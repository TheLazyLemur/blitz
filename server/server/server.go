package server

import (
	"io"
	"log/slog"
	"net"
	"os"

	"github.com/TheLazyLemur/blitz/protocol"
	"github.com/TheLazyLemur/blitz/server/ops"
)

type Config struct {
	listenAddr string
}

type Server struct {
	Config
	ln net.Listener
}

func NewServer(cfg Config) *Server {
	if len(cfg.listenAddr) <= 0 {
		cfg.listenAddr = "127.0.0.1:8080"
	}

    portFromEnv := os.Getenv("PORT")
    if len(portFromEnv) > 0 {
        cfg.listenAddr = ":" + portFromEnv
    }

	return &Server{
		Config: cfg,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}

    s.ln = listener
    defer s.ln.Close()

    slog.Info("Listenting", "port", s.listenAddr)

    return s.AcceptLoop()
}

func (s *Server) AcceptLoop() error {
    for {
        conn, err := s.ln.Accept()
        if err != nil {
			slog.Error("Error accepting connection:", "error", err.Error())
			continue
        }

        go s.handleConnection(conn)
    }
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		request, err := protocol.ParseRequest(conn)
		if err != nil {
			if err == io.EOF {
				slog.Info("Connection closed by client")
				break
			}

			slog.Error("Something went wrong", "error", err.Error())
			os.Exit(1)
		}

		go s.handleCommand(conn, request)
	}

	slog.Info("connection closed")
}

func (s *Server) handleCommand(w io.Writer, request protocol.Command) {
	var result protocol.Response

	switch request.Command {
	case "set":
		ops.HandleSet(request)
		result = protocol.NewResult(0, []byte(""))
	case "get":
		value := ops.HandleGet(request)
		result = protocol.NewResult(0, value)
	}

	protocol.SendResponse(w, result)
}
