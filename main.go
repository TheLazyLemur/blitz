package main

import (
	"encoding/binary"
	"io"
	"log/slog"
	"net"
)

var (
	port = ":8080"
)

func main() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer ln.Close()

	accteptConnection(ln)
}

// acceptConnection continuously accepts incoming connections and handles them.
func accteptConnection(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			slog.Warn("Failed to accept connection...")
			continue
		}

		go func() {
			defer conn.Close()
			handleNetworkCommunication(conn, conn)
		}()
	}
}

// handleNetworkCommunication reads a command from the incoming connection,
// validates it, and performs the corresponding action.
func handleNetworkCommunication(r io.Reader, w io.Writer) error {
	var command Command
	if err := binary.Read(r, binary.LittleEndian, &command); err != nil {
		return handleError(err)
	}

	switch command {
	case Set:
		handleSet(r, w)
	case Get:
		handleGet(r, w)
	default:
		_ = binary.Write(w, binary.LittleEndian, InvalidCommand)
		return handleError(ErrInvalidCommand)
	}

	return nil
}

// handleError handles errors returned by various functions and
// determines the appropriate action to take.
func handleError(err error) error {
	if err == io.EOF {
		return ErrConnectionClosedByClient
	}

	return err
}
