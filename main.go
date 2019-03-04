package main

import (
	"os"

	"github.com/zhanghhh369/PBFT-Blockchain/network"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID)

	server.Start()
}
