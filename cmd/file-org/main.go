package main

import (
	"log"
	"os"

	fileorg "github.com/AYehia0/file-org"
	//fileorg "github.com/AYehia0/file-org"
)

func main() {
	// the path where the files are located

	if len(os.Args) != 3 {
		log.Fatal("The target path which contains all the files and the location to put files in, are need!")
	}
	paths := os.Args

	fileorg.Run(paths[1], paths[2])
}
