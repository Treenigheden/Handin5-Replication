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
	leaderNode     int32
	client         pb.ServerNodeClient
	connectedNodes []pb.ServerNodeClient
	isLeaderNode   bool //might not be needed since we have the leadernode variable, but could improve readability
	timestamp      int32
}

type Server struct {
	pb.UnimplementedNodeServiceServer
	node NodeInfo
}

var connectedNodesMapPort = make(map[int32]NodeInfo)

var connectedNodesMapClient = make(map[pb.ServerNodeClient]NodeInfo)

func (s *Server) Bid(amount int32) (success bool) {

}

func (s *Server) Result() (amount int32, auctionOver bool) {

}

func (s *Server) AnnounceConnection(port int32, timestamp int32) (success bool) {

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

func (s *Server) RequestLederPosition(port int32) {

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
		_, err1 := node.IExist(context.Background(), &pb.Confirmation{})
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
	pb.RegisterNodeServiceServer(grpcServer, &server)

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

	go server.RequestLederPosition(node.port)

	select {}
}
