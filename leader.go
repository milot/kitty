package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"
)

// ClusterNode represents node metadata
type ClusterNode struct {
	ID        int    `json:"id"`
	IPAddress string `json:"ipAddress"`
	Port      string `json:"port"`
}

// Message represents a format for a Request/Response to for adding a node to a cluster
type Message struct {
	From    ClusterNode `json:"from"`
	To      ClusterNode `json:"to"`
	Message string      `json:"message"`
}

func main() {
	port := flag.String("port", "8001", "Port to run this node on, default 8001")

	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	currentNodeID := rand.Intn(2048)
	currentNodeIP, _ := net.InterfaceAddrs()

	currentNode := ClusterNode{
		ID:        currentNodeID,
		IPAddress: currentNodeIP[0].String(),
		Port:      *port,
	}

	fmt.Printf("Leader Node Started. Followers can connect to %v:%v\n", currentNode.IPAddress, currentNode.Port)
	startLeaderNode(currentNode)

}

func startLeaderNode(currentNode ClusterNode) {
	listener, err := net.Listen("tcp", fmt.Sprint(":"+currentNode.Port))

	if err != nil {
		fmt.Println("Port is already in use!")
	} else {
		for {
			inboundConnection, err := listener.Accept()

			if err != nil {
				if _, success := err.(net.Error); success {
					fmt.Println("Error on listen", currentNode.ID)
					return
				}
			} else {
				var requestMessage Message
				json.NewDecoder(inboundConnection).Decode(&requestMessage)
				fmt.Println(requestMessage.Message)
				inboundConnection.Close()
			}
		}
	}
}
