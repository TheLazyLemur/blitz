package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"sync"

	"github.com/TheLazyLemur/blitz/protocol"
	"github.com/TheLazyLemur/blitz/types"
)

func main() {
	serverAddr := "localhost:8080"

	c := NewClient(serverAddr)
	defer c.Close()

	i := 0
	for {
		i++

		c.Set("test"+fmt.Sprintf("%d", i), "its working"+fmt.Sprintf("%d", i))
		c.Get("test" + fmt.Sprintf("%d", i))

		if i == 100000 {
			break
		}
	}
}

type Client struct {
	conn net.Conn
    lock sync.Mutex
}

func NewClient(serverAddr string) *Client {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	return &Client{
		conn: conn,
        lock: sync.Mutex{},
	}
}

func (c *Client) Close() {
	defer c.conn.Close()
}

func (c *Client) Get(k string) types.GetResponse {
    c.lock.Lock()
    defer c.lock.Unlock()

	msg := types.GetRequest{
		BaseMessage: protocol.BaseMessage{
			JsonRPC: "2.0",
			Method:  types.Get,
		},
		Key: k,
	}

	bs, _ := protocol.EncodeMessage(msg)

	scanner := bufio.NewScanner(c.conn)
	scanner.Split(protocol.ScanLines)

	_, err := fmt.Fprintf(c.conn, bs)
	if err != nil {
		if err == io.EOF {
			slog.Info("Connection closed by server")
		}

		return types.GetResponse{}
	}

	response, err := scanGetResponse(scanner)
	if err != nil {
		panic(err)
	}

	slog.Info("server response", "version", response.JsonRPC, "method", response.Method, "code", response.ResponseCode, "value", response.Value)
	return response
}

func (c *Client) Set(k, v string) types.SetResponse {
    c.lock.Lock()
    defer c.lock.Unlock()

	msg := types.SetRequest{
		BaseMessage: protocol.BaseMessage{
			JsonRPC: "2.0",
			Method:  types.Set,
		},
		Key:   k,
		Value: v,
	}

	bs, _ := protocol.EncodeMessage(msg)

	scanner := bufio.NewScanner(c.conn)
	scanner.Split(protocol.ScanLines)

	_, err := fmt.Fprintf(c.conn, bs)
	if err != nil {
		if err == io.EOF {
			slog.Info("Connection closed by server")
		}

		return types.SetResponse{}
	}

	response, err := scanSetResponse(scanner)
	if err != nil {
		panic(err)
	}

	slog.Info("server response", "version", response.JsonRPC, "method", response.Method, "code", response.ResponseCode)
	return response

}

func scanGetResponse(scanner *bufio.Scanner) (types.GetResponse, error) {
	for scanner.Scan() {
		msg := scanner.Bytes()

		_, contents, err := protocol.DecodeMessage(msg)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		result := &types.GetResponse{}
		err = json.Unmarshal(contents, result)
		if err != nil {
			panic(err.Error())
		}

		return *result, scanner.Err()
	}

	return types.GetResponse{}, nil
}

func scanSetResponse(scanner *bufio.Scanner) (types.SetResponse, error) {
	for scanner.Scan() {
		msg := scanner.Bytes()

		_, contents, err := protocol.DecodeMessage(msg)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		result := &types.SetResponse{}
		err = json.Unmarshal(contents, result)
		if err != nil {
			panic(err.Error())
		}

		return *result, scanner.Err()
	}

	return types.SetResponse{}, nil
}
