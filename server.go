package main

import (
	"encoding/binary"
	"io"
	"log/slog"
	"net"
	"os"
)

type Server struct {
	listenAddrr string
	ln          net.Listener
}

func NewServer(listenAddrr string) *Server {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return &Server{
		listenAddrr: listenAddrr,
		ln:          ln,
	}
}

func (s *Server) StartServer() error {
	defer s.ln.Close()
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Warn("Failed to accept connection...")
			continue
		}

		go func() {
			defer conn.Close()
			s.handleNetworkCommunication(conn, conn)
		}()
	}
}

// handleNetworkCommunication reads a command from the incoming connection,
// validates it, and performs the corresponding action.
func (s *Server) handleNetworkCommunication(r io.Reader, w io.Writer) error {
	var command Command
	if err := binary.Read(r, binary.LittleEndian, &command); err != nil {
		return handleError(err)
	}

	switch command {
	case Set:
		if err := hnds.handleSet(r, w); err != nil {
			return handleError(err)
		}
	case Get:
		if err := hnds.handleGet(r, w); err != nil {
			return handleError(err)
		}
	default:
		_ = binary.Write(w, binary.LittleEndian, InvalidCommand)
		return handleError(ErrInvalidCommand)
	}

	return nil
}
