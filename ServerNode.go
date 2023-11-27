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

var connectedNodesMapPort = make(map[int32]NodeInfo)

var connectedNodesMapClient = make(map[pb.ServerNodeClient]NodeInfo)
var highestBid = -1
var highestBidderID = -1
var auctionIsRunning = false

func (s *Server) Bid(ctx context.Context, input *pb.BidInput) (*pb.Confirmation, error) {
	if !auctionIsRunning {
		fmt.Println("Someone tried to bed while the auction was not running!")
		return &pb.Confirmation{Success: false}, nil
	} else if input.Bid > int32(highestBid) {
		highestBid = int(input.Bid)
		highestBidderID = int(input.Port)
		fmt.Printf("New highest bid: %v New highest bidder: %v\n", highestBid, highestBidderID)
		return &pb.Confirmation{Success: true}, nil
	} else {
		return &pb.Confirmation{Success: false}, nil
	}
}

func (s *Server) Result(ctx context.Context, _ *pb.Empty) (*pb.Outcome, error) {
	if auctionIsRunning {
		return &pb.Outcome{Amount: int32(highestBid), AuctionOver: false, Winner: int32(highestBidderID)}, nil
	} else {
		return &pb.Outcome{Amount: int32(highestBid), AuctionOver: true, Winner: int32(highestBidderID)}, nil
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
	//we have establised a new connection to the new node. We add it to the list of node connections
	node := pb.NewServerNodeClient(conn)
	//Add the node we have connected to our list of nodes in the system.
	//We also maintain a map which lets us find the node from its NodeID.
	nodeInfo := NodeInfo{port: announcement.NodeID, client: node}
	s.node.connectedNodes = append(s.node.connectedNodes, nodeInfo.client)
	connectedNodesMapPort[announcement.NodeID] = nodeInfo
	connectedNodesMapClient[node] = nodeInfo
	//We send back a confirmation message to indicate that the connection was esablished
	return &pb.Confirmation{}, nil
}

func (s *Server) AnnounceUpdate(ctx context.Context, announcement *pb.UpdateAnnouncement) (*pb.Confirmation, error) {
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
	s.node.leader = node.client

	return &pb.Empty{}, nil
}

func (s *Server) RequestLederPosition() {
	var isLeaderCandidate = true
	for _, connectedNode := range s.node.connectedNodes {
		var response, err = connectedNode.RequestLeadership(context.Background(), &pb.AccessRequest{NodeID: s.node.port, Timestamp: s.node.timestamp})
		if err != nil {
			log.Fatalf("OH NO, AN ERROR OCCURED WHILE ASKING A NODE TO BE A LEADER")
		}
		if !response.Granted {
			isLeaderCandidate = false
			break
		}
	}

	if isLeaderCandidate {
		fmt.Println("I am new leader")
		s.node.isLeaderNode = true
		s.node.leader = s.node.client
		s.node.timestamp++
		s.AnnounceLeadership()
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
		connectedNodes = append(connectedNodes, node)                                                                                     //OBS OBS OBS - why do we have two of these? Can we just have the connected nodes on the server? Why is it on the server anyway?
		nodeInfo := &NodeInfo{port: int32(port), client: node, connectedNodes: s.node.connectedNodes, timestamp: int32(s.node.timestamp)} //OBS OBS OBS IS THIS RIGHT REGARDING TIMESTAMP?
		connectedNodesMapPort[int32(port)] = *nodeInfo
		connectedNodesMapClient[node] = *nodeInfo

		//Then we send an announcement to inform the node
		//in order to inform it that we have connected to it (and that it should connect to this node in return.)
		_, err = node.AnnounceConnection(context.Background(), &pb.ConnectionAnnouncement{NodeID: int32(s.node.port)})
		if err != nil {
			log.Fatalf("Oh no! The node sent an announcement of a new connection but did not recieve a confirmation in return. Error: %v", err)
		}
		//fmt.Println("SENDING THE ANNOUNCEMENT SEEMS TO HAVE GONE OK?")

	}
}

func (s *Server) cli_interface() {
	for {
		fmt.Print(" > ")
		var input string
		fmt.Scanln(&input)
		if input == "result" {
			result, err := s.node.leader.Result(context.Background(), &pb.Empty{})
			if err != nil {
				fmt.Println("ERROR OCCURED! TRY AGAIN.")
			}
			if result.AuctionOver { //OBS Why capital first letter? Is this correct?
				fmt.Println("The auction is over. The winning bid was: ", result.Amount)
			} else {
				fmt.Println("The auction is ongoing. The current highest bid is:", result.Amount)
			}
		} else if input == "start" {
			if s.node.isLeaderNode {
				auctionIsRunning = true
				fmt.Println("Starting new auction ...")
			} else {
				fmt.Println("Only the auction leader can start and end auctions ...")
			}
		} else if input == "end" && s.node.isLeaderNode {
			result, err := s.node.leader.Result(context.Background(), &pb.Empty{})
			if err != nil {
				fmt.Println("ERROR OCCURED! TRY AGAIN.")
			}
			auctionIsRunning = false
			fmt.Println("Ending auction! The winning bid was: " + strconv.Itoa(int(result.Amount)))
		} else {
			inputInt, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("INVALID INPUT! TRY AGAIN.")
			} else {
				s.node.leader.Bid(context.Background(), &pb.BidInput{Bid: int32(inputInt), Port: s.node.port})
			}
		}

	}
}

func main() {
	timestamp := 0
	connectedNodes := []pb.ServerNodeClient{}

	//finds a port to listen on
	standardPort := 8000
	port, err := FindAnAvailablePort(standardPort)
	fmt.Printf("MY PORT IS: %v\n ", port)
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
	connectedNodesMapPort[int32(port)] = *node
	connectedNodesMapClient[thisNodeClient] = *node
	server.node = *node

	server.EstablishConnectionToAllOtherNodes(standardPort, port, transportCreds, connectedNodes)
	log.Printf("The number of connected nodes is %v", len(server.node.connectedNodes))

	go server.RequestLederPosition()
	go server.cli_interface()

	select {}
}
