package main

import (
	"log"

	"github.com/TheLazyLemur/blitz/server/server"
)

func main() {
    s := server.NewServer(server.Config{})
    log.Fatal(s.Start())
}
