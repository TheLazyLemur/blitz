package main

import (
	"log"
)

var (
	port = ":8080"
	hnds *handlers
)

func main() {
	hnds = &handlers{}

	s := NewServer(port)
	log.Fatal(s.StartServer())
}
