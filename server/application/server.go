package application

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"net"

	"github.com/TheLazyLemur/blitz/protocol"
	"github.com/TheLazyLemur/blitz/server/store"
	"github.com/TheLazyLemur/blitz/types"
)

type ServerOptions struct {
	ServerAddr string
	Store      store.Storer
}

type Server struct {
	ServerOptions
	ln    net.Listener
}

func NewServer(opts ServerOptions) *Server {
	return &Server{
		ServerOptions: opts,
	}
}

func (s *Server) Connect() error {
	listener, err := net.Listen("tcp", s.ServerAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return err
	}

	s.ln = listener

	fmt.Println("Server listening on port", s.ServerAddr)
	return nil
}

func (s *Server) Start() error {
	defer s.ln.Close()

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			return err
		}

		go s.handleConn(conn, conn)
	}
}

func (s *Server) handleConn(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Split(protocol.ScanLines)

	for scanner.Scan() {
		msg := scanner.Bytes()

		method, contents, err := protocol.DecodeMessage(msg)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		s.handleMessage(w, method, contents)

	}

	if scanner.Err() != nil {
		if scanner.Err() == io.EOF {
			slog.Info("Connection closed by client")
		} else {
			slog.Error("Scanner error:", scanner.Err())
		}
	} else {
		fmt.Println("Connection closed by client")
	}
}

func (s *Server) handleMessage(writer io.Writer, method string, contents []byte) {
	switch method {
	case types.Set:
		s.handleSet(writer, contents)
	case types.Get:
		s.handleGet(writer, contents)
	default:
		panic("not implemented")
	}
}
