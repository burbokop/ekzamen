package main

import (
	"os"

	"github.com/burbokop/ekzamen/server/tree"
)

func main() {
	binTree := &tree.BinaryTree{}
	binTree.insert(100.0)

	print(os.Stdout, binTree.root, 0, 'M')

	// Create the server.
	//if err := StartServer(); err == nil {
	//	log.Println("Starting chat server...")
	//} else {
	//	log.Fatalf("Cannot initialize banking server: %s", err)
	//}
}
