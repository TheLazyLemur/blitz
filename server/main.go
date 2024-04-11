package main

import (
	"log"

	"github.com/TheLazyLemur/blitz/server/application"
	"github.com/TheLazyLemur/blitz/server/store"
)

func main() {
	port := ":8080"
    s := application.NewServer(application.ServerOptions{
        Store: store.NewMemStore(),
        ServerAddr: port,
    })

    s.Store.Set("Hello", "world")

    if err := s.Connect(); err != nil {
        log.Fatal(err)
    }

    log.Fatal(s.Start())
}
