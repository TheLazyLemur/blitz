package main

import (
	"log"
)

var (
	port = ":8080"
)

func main() {
	hnds := &handlersImpl{}

	s := NewServer(port, hnds)
	log.Fatal(s.StartServer())
}
