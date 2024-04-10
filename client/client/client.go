package client

import (
	"log/slog"
	"net"
	"sync"

	"github.com/TheLazyLemur/blitz/protocol"
)

type Client struct {
	serverAddr string
	conn       net.Conn
	lock       sync.Mutex
}

func NewClient(severAddr string) *Client {
	return &Client{
		serverAddr: severAddr,
		lock:       sync.Mutex{},
	}
}

func (c *Client) Connect() *Client {
	conn, err := net.Dial("tcp", c.serverAddr)
	if err != nil {
		slog.Error("Error connecting", "error", err.Error())
		return nil
	}

	c.conn = conn

	return c
}

func (c *Client) MakeRequest(cmd protocol.Command) (protocol.Response, error) {
    c.lock.Lock()
    defer c.lock.Unlock()

	if err := protocol.SendCommand(c.conn, cmd); err != nil {
		return protocol.Response{}, err
	}

	result, err := protocol.ReadResponse(c.conn)
	if err != nil {
		return protocol.Response{}, err
	}

	return result, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
