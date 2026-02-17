package main

import (
	"fmt"
	"log"

	"github.com/ConnorBrightman/ezauth/internal/config"
)

func runInit() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ ezauth initialized successfully.")
	fmt.Println("Next step: run `ezauth start`")
}
