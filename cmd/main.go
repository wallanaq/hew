package main

import (
	"log"

	"github.com/wallanaq/hew/cmd/root"
)

func main() {
	rootCmd := root.NewRootCommand()

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
