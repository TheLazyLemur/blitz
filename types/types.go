package types

import "github.com/TheLazyLemur/blitz/protocol"

type Response struct {
	protocol.BaseMessage
	ResponseCode int
}

type SetRequest struct {
	protocol.BaseMessage
	Key   string
	Value string
}

type SetResponse struct {
	Response
}

type GetRequest struct {
	protocol.BaseMessage
	Key   string
}

type GetResponse struct {
	Response
	Value string
}

const (
	Set = "set"
	Get = "get"
)
