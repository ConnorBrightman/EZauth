package main

import "fmt"

func printHelp() {
	fmt.Println("ezauth - lightweight local authentication server")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  ezauth init       Initialize config.yaml")
	fmt.Println("  ezauth start      Start authentication server")
	fmt.Println("  ezauth version    Show version")
}
