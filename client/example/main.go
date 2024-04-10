package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"

	"github.com/TheLazyLemur/blitz/client/client"
	"github.com/TheLazyLemur/blitz/protocol"
)

func main() {
	serverAddr := "127.0.0.1:8080"

	c := client.NewClient(serverAddr).Connect()
	defer c.Close()

	slog.Info("Connected to", "addr", serverAddr)

	for {
		fmt.Println("command (get, set)")
		var inputCmd string
		_, _ = fmt.Scanln(&inputCmd)

		fmt.Println("cache key")
		var inputKey string
		_, _ = fmt.Scanln(&inputKey)

		var inputValue string
		if inputCmd == "set" {
			fmt.Println("cache value")
			_, _ = fmt.Scanln(&inputValue)
		}

		cmd := protocol.NewCommand(inputCmd, inputKey, inputValue)

		result, err := c.MakeRequest(cmd)
		if err != nil {
			if err == io.EOF {
				slog.Info("Connection closed by server")
				break
			}

			log.Fatal(err)
		}

		_ = result

	}

	slog.Info("connection closed")
}
