package main

import (
	"github.com/dewmal/rakun/cmd"
	"log"
)

func main() {
	log.Printf("Application - %s begin to start", VERSION)
	cmd.Execute()
}
