/*
idea for gennerel code structure:
create nodes
*/
package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

type NodeInfo struct {
	port int32
	//client                        pb.NodeServiceClient
	//connectedNodes                []pb.NodeServiceClient
	isLeaderNode bool
	timestamp    int32
	//localQueue                    []pb.AccessRequest
}

type Server struct {
	//pb.UnimplementedNodeServiceServer
	node NodeInfo
}

var connectedNodesMapPort = make(map[int32]NodeInfo)

//var connectedNodesMapClient = make(map[pb.NodeServiceClient]NodeInfo)

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

func (s *Server) AwaitClients(port int32) {
	//not implimentet yet
}

func main() {
	timestamp := 0
	//connectedNodes := []pb.NodeServiceClient{}

	standardPort := 8000
	port, err := FindAnAvailablePort(standardPort)
	fmt.Printf("THE PORT IS: %v\n ", port)
	if err != nil {
		log.Fatalf("Oh no! Failed to find a port")
	}

	// Create a gRPC server
	//grpcServer := grpc.NewServer()
	server := Server{}
	// Register your gRPC service with the server
	//pb.RegisterNodeServiceServer(grpcServer, &server)

	//node := &NodeInfo{port: int32(port), client: thisNodeClient, connectedNodes: connectedNodes, timestamp: int32(timestamp)}
	node := &NodeInfo{port: int32(port), timestamp: int32(timestamp)}
	server.node = *node

	//server.EstablishConnectionToAllOtherNodes()

	go server.AwaitClients(node.port)

	select {}
}
