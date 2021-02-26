package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
	"kitty/node"
)

func main() {
	clusterIP := flag.String("connectTo", "127.0.0.1:8001", "IP Address for nodes to connect")
	port := flag.String("port", "8001", "Port to run this node on, default 8001")

	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	currentNodeID := rand.Intn(2048)
	currentNodeIP, _ := net.InterfaceAddrs()

	currentNode := node.Node{
		ID:        currentNodeID,
		IPAddress: currentNodeIP[0].String(),
		Port:      *port,
	}

	destinationNode := node.Node{
		ID:        -1,
		IPAddress: strings.Split(*clusterIP, ":")[0],
		Port:      strings.Split(*clusterIP, ":")[1],
	}

	fmt.Printf("Started follower node %v IP: %v - Port: %v\n", currentNodeID, currentNodeIP[0].String(), *port)

	for {
		time.Sleep(time.Second * 2)
		startFollowerNode(currentNode, destinationNode)
	}

}

func constructMessage(source node.Node, dest node.Node, message string) node.Message {
	return node.Message {
		From: node.Node {
			ID:        source.ID,
			IPAddress: source.IPAddress,
			Port:      source.Port,
		},
		To: node.Node {
			ID:        dest.ID,
			IPAddress: dest.IPAddress,
			Port:      dest.Port,
		},
		Message: message,
	}
}

func startFollowerNode(currentNode node.Node, destinationNode node.Node) {
	outboundConnection, err := net.DialTimeout("tcp", destinationNode.IPAddress+":"+destinationNode.Port, time.Duration(10)*time.Second)

	if err != nil {
		if _, success := err.(net.Error); success {
			fmt.Println("Could not connect to the cluster. Retrying...", currentNode.ID)
		}
	} else {
		text := fmt.Sprintf("Follower Node %v with IP %v:%v is following you.", currentNode.ID, currentNode.IPAddress, currentNode.Port)
		requestMessage := constructMessage(currentNode, destinationNode, text)
		json.NewEncoder(outboundConnection).Encode(&requestMessage)

		decoder := json.NewDecoder(outboundConnection)
		var responseMessage node.Message
		decoder.Decode(&responseMessage)
		fmt.Printf("Message sent to the leader %v:%v\n", destinationNode.IPAddress, destinationNode.Port)
	}
}
