package node


// ClusterNode represents node metadata
type Node struct {
	ID        int    `json:"id"`
	IPAddress string `json:"ipAddress"`
	Port      string `json:"port"`
}

// Message represents a format for a Request/Response to for adding a node to a cluster
type Message struct {
	From    Node 		`json:"from"`
	To      Node 		`json:"to"`
	Message string      `json:"message"`
}