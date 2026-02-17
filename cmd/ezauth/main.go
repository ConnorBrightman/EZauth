package main

import (
	"fmt"
	"log"
	"os"
)

var Version = "dev"

func main() {
	Execute()
}

func Execute() {
	// Banner
	log.Println(`
 _____ _____         _   _   
|   __|__   |___ _ _| |_| |_ 
|   __|   __| .'| | |  _|   |
|_____|_____|__,|___|_| |_|_|              
Authentication made EZ          
`)

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {

	case "init":
		runInit()

	case "start":
		runStart()

	case "version", "--version", "-v":
		fmt.Println("ezauth version:", Version)

	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printHelp()
		os.Exit(1)
	}
}
