package main

import (
	"log"

	"github.com/noi/go-snippets/uds/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("ERROR: %+v\n", err)
	}
}
