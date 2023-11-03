package main

import (
	"github.com/csyezheng/a2fa/cmd/commands"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	err := commands.Execute(os.Args[1:])
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
