/*
idea for gennerel code structure:
create nodes
*/
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	pb "homework5/gRPC"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type NodeInfo struct {
	port           int32
	leader         pb.ServerNodeClient
	client         pb.ServerNodeClient
	connectedNodes []pb.ServerNodeClient
	isLeaderNode   bool //might not be needed since we have the leadernode variable, but could improve readability
	timestamp      int32
}

type Server struct {
	pb.UnimplementedServerNodeServer
	node NodeInfo
}

var connectedNodesMapPort = make(map[int32]pb.ServerNodeClient)
var connectedNodesMapClient = make(map[pb.ServerNodeClient]int32)
var highestBid = -1
var highestBidderID = -1
var auctionIsRunning = false
var timerIsRunning = false

func (s *Server) Bid(ctx context.Context, input *pb.BidInput) (*pb.Confirmation, error) {
	s.node.timestamp++
	//fmt.Println("The timestamp is ", s.node.timestamp)
	if !auctionIsRunning {
		fmt.Println(" * * * Someone tried to bid while the auction was not running! * * * ")
		return &pb.Confirmation{Success: false, Timestamp: s.node.timestamp - 1}, nil
	} else if input.Bid > int32(highestBid) {
		highestBid = int(input.Bid)
		highestBidderID = int(input.Port)
		fmt.Printf(" * * * New highest bid: %v * * * \n * * * Highest bidder: %v * * * \n", highestBid, highestBidderID)
		return &pb.Confirmation{Success: true, Timestamp: s.node.timestamp - 1}, nil
	} else {
		return &pb.Confirmation{Success: false, Timestamp: s.node.timestamp - 1}, nil
	}
}

func (s *Server) Result(ctx context.Context, _ *pb.Empty) (*pb.Outcome, error) {
	s.node.timestamp++
	//fmt.Println("The timestamp is ", s.node.timestamp)

	if auctionIsRunning {
		return &pb.Outcome{Amount: int32(highestBid), AuctionOver: false, Winner: int32(highestBidderID), Timestamp: s.node.timestamp - 1}, nil
	} else {
		return &pb.Outcome{Amount: int32(highestBid), AuctionOver: true, Winner: int32(highestBidderID), Timestamp: s.node.timestamp - 1}, nil
	}
}

func (s *Server) AnnounceConnection(ctx context.Context, announcement *pb.ConnectionAnnouncement) (*pb.Confirmation, error) {
	//We have recieved a connection announcement, which means that a new node has established a connection to this client.
	//We must also establish a connection to this client in return. We have the information we need from the ConnectionAnnouncement
	transportCreds := insecure.NewCredentials()
	//Establish a grpc connection to the other node using addres and transport credentials
	address := ":" + strconv.Itoa(int(announcement.NodeID))
	fmt.Println("NEW NODE JOINED.")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(transportCreds))
	if err != nil {
		log.Fatalf("Failed to connect  ... : %v\n", err)
	}
	//we have establised a new connection to the new node.
	node := pb.NewServerNodeClient(conn)
	//We add the node we have connected to our list of nodes in the system.
	//We also maintain a map which lets us find the node from its NodeID.
	s.node.connectedNodes = append(s.node.connectedNodes, node)
	connectedNodesMapPort[announcement.NodeID] = node
	connectedNodesMapClient[node] = announcement.NodeID
	//We update the node with the newest information
	node.AnnounceUpdate(context.Background(), &pb.UpdateAnnouncement{HighestBid: int32(highestBid), HighestBidder: int32(highestBidderID), AuctionIsOngoing: auctionIsRunning, Timestamp: s.node.timestamp})
	//We send back a confirmation message to indicate that the connection was esablished
	return &pb.Confirmation{}, nil
}

func (s *Server) AnnounceUpdate(ctx context.Context, announcement *pb.UpdateAnnouncement) (*pb.Confirmation, error) {
	//We need to keep nodes updated in case they need to take over as leader.
	//This means updating every node each time something new happens.
	highestBid = int(announcement.HighestBid)
	highestBidderID = int(announcement.HighestBidder)
	auctionIsRunning = announcement.AuctionIsOngoing
	return &pb.Confirmation{}, nil //OBSOBSOBSOBS
}

func (s *Server) RequestLeadership(ctx context.Context, request *pb.AccessRequest) (*pb.AccessRequestResponse, error) {
	//When this method is called by a remote node, it means that the current leader is dead and there is an election for leadership in process. '
	//First we compare the timestamp of this node with the timestamp of the incoming request.
	//If the incoming timestamp is greater (OBS) than the local timestamp, we know that we are not going to win an election. We don't bother sending out requests. We send a response.
	//If the incoming timestamp is less than the local timestamp, we know that the remote node is out. We have a chance ourselves, however! We send out requests. We send a response.
	if request.Timestamp >= s.node.timestamp {
		return &pb.AccessRequestResponse{Granted: true, Timestamp: s.node.timestamp}, nil
	} else {
		go s.RequestLederPosition()
		return &pb.AccessRequestResponse{Granted: false, Timestamp: s.node.timestamp}, nil
	}
}

func (s *Server) IExist(ctx context.Context, _ *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil

}

func (s *Server) IAmLeader(ctx context.Context, anouncement *pb.ConnectionAnnouncement) (*pb.Empty, error) {
	//When this method is called by a remote node, it means that that node is the new leader.
	//We update the leader node and send back a confirmation message (empty)
	var node = connectedNodesMapPort[anouncement.NodeID]
	s.node.leader = node

	return &pb.Empty{}, nil
}

func (s *Server) RequestLederPosition() {
	var isLeaderCandidate = true
	for _, connectedNode := range s.node.connectedNodes {
		var response, err = connectedNode.RequestLeadership(context.Background(), &pb.AccessRequest{NodeID: s.node.port, Timestamp: s.node.timestamp})
		if err != nil {
			log.Println("A node is down ... this may or may not be expected.")
		} else if !response.Granted {
			isLeaderCandidate = false
			break
		}
	}

	if isLeaderCandidate {
		fmt.Println(" * * * This node is the leader * * * ")
		if !timerIsRunning {
			if auctionIsRunning {
				fmt.Println("Restarting auction timer!")
				go RunTimerOneMin()
			}
		}

		s.node.isLeaderNode = true
		s.node.leader = s.node.client
		s.AnnounceLeadership()
		s.node.timestamp++
		//fmt.Println("The timestamp is ", s.node.timestamp)
	}

}

func (s *Server) AnnounceLeadership() {
	//To announce this node's leadership we cycle through connected nodes and call IAmLeader
	for _, connectedNode := range s.node.connectedNodes {
		connectedNode.IAmLeader(context.Background(), &pb.ConnectionAnnouncement{NodeID: s.node.port, Timestamp: s.node.timestamp})
	}

}

func FindAnAvailablePort(standardPort int) (int, error) {
	for port := standardPort; port < standardPort+100; port++ {
		addr := "localhost:" + strconv.Itoa(port)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			//The port is in use, increment and try the next one
			continue
		}
		//if no error, the port is free. Return the port.
		listener.Close()
		return port, nil
	}
	return 0, fmt.Errorf("no free port found in the range")
}

func (s *Server) EstablishConnectionToAllOtherNodes(standardPort int, thisPort int, transportCreds credentials.TransportCredentials, connectedNodes []pb.ServerNodeClient) {
	//We cycle through the available ports in order to find the other nodes in the system and establish connections.
	for port := standardPort; port < standardPort+100; port++ {
		if port == thisPort {
			continue
		}
		address := "localhost:" + strconv.Itoa(port)
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(transportCreds))
		if err != nil {
			continue
		}

		//We make a node with the connection to check if there is anything there
		node := pb.NewServerNodeClient(conn)
		_, err1 := node.IExist(context.Background(), &pb.Empty{})
		if err1 != nil {
			//There is no node on this port. We move onto the next and try again.
			continue
		}
		//We have established a connection to existing port!
		//First we add the node to our own list of connected nodes as well as the relevant maps.....
		s.node.connectedNodes = append(s.node.connectedNodes, node)
		//connectedNodes = append(connectedNodes, node)

		connectedNodesMapPort[int32(port)] = node
		connectedNodesMapClient[node] = int32(port)

		//Then we send an announcement to inform the node
		//in order to inform it that we have connected to it (and that it should connect to this node in return.)
		_, err = node.AnnounceConnection(context.Background(), &pb.ConnectionAnnouncement{NodeID: int32(s.node.port)})
		if err != nil {
			log.Fatalf("Oh no! The node sent an announcement of a new connection but did not recieve a confirmation in return. Error: %v", err)
		}
		//fmt.Println("SENDING THE ANNOUNCEMENT SEEMS TO HAVE GONE OK?")

	}
}

func RunTimerOneMin() {
	timerIsRunning = true
	for i := 0; i < 60; i++ {
		time.Sleep(time.Duration(1) * time.Second)
	}

	auctionIsRunning = false
	timerIsRunning = false
	fmt.Println("TIME IS UP!")
}

func (s *Server) cli_interface() {

	fmt.Println("Type 'result' to see details about the auction. \nType any number to bid in the action.")
	if s.node.isLeaderNode {
		fmt.Println("Start and end auctions by typing 'start' or 'end' ")
	}
	for {
		//fmt.Print(" > ")
		//fmt.Println("Current timestamp: ", s.node.timestamp)
		var input string
		fmt.Scanln(&input)
		if input == "result" {
			result, err := s.node.leader.Result(context.Background(), &pb.Empty{})
			if err != nil {
				//If we get an error back from the leader we assume that the leader is dead. We request leadership.
				s.RequestLederPosition()
				//Once a new leader has been determined we repeat the result method call:
				time.Sleep(time.Millisecond * 100) //We sleep to allow the new leader to be elected before continuing.
				result, _ = s.node.leader.Result(context.Background(), &pb.Empty{})
			}
			auctionIsRunning = !result.AuctionOver
			highestBid = int(result.Amount)
			highestBidderID = int(result.Winner)
			s.updateTimestamp(result.Timestamp) //TIMESTAMP
			if result.AuctionOver {
				fmt.Println("There is no ongoing auction. ")
				if result.Amount == -1 {
					fmt.Println("The were no bids.")
				} else {
					fmt.Printf("The winning bidder was: %v with a bid of: %v. \n", result.Winner, result.Amount)
				}
			} else {
				fmt.Println("The auction is ongoing. ")
				if result.Amount == -1 {
					fmt.Println("There are no bids yet.")
				} else {
					fmt.Printf("The current highest bidder is: %v with a bid of: %v. \n", result.Winner, result.Amount)
				}
			}
		} else if input == "start" {
			if s.node.isLeaderNode {
				auctionIsRunning = true
				highestBid = -1
				highestBidderID = -1
				fmt.Println("New auction started ...")
				s.node.timestamp++ //TIMESTAMP
				//fmt.Println("The timestamp is ", s.node.timestamp)
				//s.updateAllNodes()

				go RunTimerOneMin()
			} else {
				fmt.Println("Invalid command. Only the auction leader can start and end auctions")
			}
		} else if input == "end" && s.node.isLeaderNode {
			result, err := s.node.leader.Result(context.Background(), &pb.Empty{})
			if err != nil {
				fmt.Println("ERROR OCCURED! TRY AGAIN.")
			}
			auctionIsRunning = false
			fmt.Println("Ending auction! The winning bid was: " + strconv.Itoa(int(result.Amount)))
			s.node.timestamp++ //TIMESTAMP
			//fmt.Println("The timestamp is ", s.node.timestamp)
			//s.updateAllNodes()
		} else {
			inputInt, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("INVALID INPUT! TRY AGAIN.")
			} else {
				outcome, _ := s.node.leader.Bid(context.Background(), &pb.BidInput{Bid: int32(inputInt), Port: s.node.port})
				auctionState, err := s.node.leader.Result(context.Background(), &pb.Empty{})

				if err != nil {
					//If we get an error back from the leader we assume that the leader is dead. We request leadership.
					s.RequestLederPosition()
					//Once a new leader has been determined we repeat the bid:
					time.Sleep(time.Millisecond * 100) //We sleep to allow the new leader to be elected before continuing.
					outcome, _ = s.node.leader.Bid(context.Background(), &pb.BidInput{Bid: int32(inputInt), Port: s.node.port})
					auctionState, _ = s.node.leader.Result(context.Background(), &pb.Empty{})
					//OBSOBSOBS ... Hmm, if the new leader fails in between becoming leader and recieving the bid, how do we handle this in a more elegant way?
				}
				if outcome.Success {
					fmt.Println("Your bid was successful.")
					auctionIsRunning = !auctionState.AuctionOver
					highestBid = inputInt
					highestBidderID = int(s.node.port)
					s.updateTimestamp(outcome.Timestamp) //TIMESTAMP

					//In order to prevent a case where the leader node has made a bid in the auction and then crashes, meaning no other node would have heard
					//of this bid unless they happened to have called result, we need to update at least one other node about the state of the auction
					//So that this node may be made leader in the described case. We find a node from the list of connected nodes (if any) and update this.
					//In this way we still avoid having to update all nodes everywhere all the time.

					if len(s.node.connectedNodes) > 1 {
						s.node.connectedNodes[len(s.node.connectedNodes)-1].AnnounceUpdate(context.Background(), &pb.UpdateAnnouncement{HighestBid: int32(highestBid), HighestBidder: int32(highestBidderID), AuctionIsOngoing: auctionIsRunning, Timestamp: s.node.timestamp})
					}

				} else {
					fmt.Println("Your bid was not successful.")
					auctionIsRunning = !auctionState.AuctionOver
					highestBid = int(auctionState.Amount)
					highestBidderID = int(auctionState.Winner)
					s.updateTimestamp(auctionState.Timestamp) //TIMESTAMP
				}
			}
		}
	}
}

func (s *Server) updateTimestamp(incomingTimestamp int32) {
	if s.node.timestamp < incomingTimestamp {
		s.node.timestamp = incomingTimestamp
	}
	s.node.timestamp++
	//fmt.Println("The timestamp is ", s.node.timestamp)
}

func main() {
	timestamp := 0
	connectedNodes := []pb.ServerNodeClient{}

	//finds a port to listen on
	standardPort := 8000
	port, err := FindAnAvailablePort(standardPort)
	fmt.Printf("MY PORT IS: %v\n\n", port)
	if err != nil {
		log.Fatalf("Oh no! Failed to find a port")
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()
	server := Server{}
	// Register your gRPC service with the server
	pb.RegisterServerNodeServer(grpcServer, &server)

	//initialize the listener on the specified port. net.Listen listens for incoming connections with tcp socket
	listen, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Could not listen at port: %d : %v", port, err)
	}
	go func() {
		// Start gRPC server in a goroutine
		err := grpcServer.Serve(listen)
		if err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	//we create insecure transport credentials (in the context of this assignment we choose not to worry about security):
	transportCreds := insecure.NewCredentials()

	//Establish a grpc connection to the other nodes using addres and transport credentials
	address := ":" + strconv.Itoa(port)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(transportCreds))
	if err != nil {
		log.Fatalf("Failed to connect  ... : %v\n", err)
	}

	//Create a grpc client instance to represent local node (this node) and add it to connected nodes.
	//!---maybe we dont need to do this for this implimetation---!
	thisNodeClient := pb.NewServerNodeClient(conn)
	connectedNodes = append(connectedNodes, thisNodeClient)

	//Genrate node
	node := &NodeInfo{port: int32(port), client: thisNodeClient, connectedNodes: connectedNodes, timestamp: int32(timestamp)}
	connectedNodesMapPort[int32(port)] = thisNodeClient
	server.node = *node

	server.EstablishConnectionToAllOtherNodes(standardPort, port, transportCreds, connectedNodes)
	//log.Printf("The number of connected nodes is %v", len(server.node.connectedNodes))

	go server.RequestLederPosition()
	time.Sleep(1000 * time.Millisecond)
	go server.cli_interface()

	select {}
}
