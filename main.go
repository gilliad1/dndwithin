package main

import (
	"fmt"
	"log"
)

func main() {
	server, serverErr := newServer()
	if serverErr != nil {
		log.Fatalf("ERROR INITIALIZING SERVER: %s", serverErr.Error())
	}
	e, err := server.GetEchoServer()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(e)
}
