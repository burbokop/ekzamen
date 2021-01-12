package main

import (
	"os"

	"github.com/burbokop/ekzamen/server/tree"
)

func main() {
	binTree := &tree.BinaryTree{}

	binTree.Insert(4.0, true).Insert(2.0, false).Insert(8.0, true).Insert(10, false)

	tree.Print(os.Stdout, binTree.Root, 0, 'M')

	tree.MultiplyTree(binTree)

	tree.Print(os.Stdout, binTree.Root, 0, 'M')

	// Create the server.
	//if err := StartServer(); err == nil {
	//	log.Println("Starting chat server...")
	//} else {
	//	log.Fatalf("Cannot initialize banking server: %s", err)
	//}
}
