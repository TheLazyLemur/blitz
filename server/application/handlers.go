package application

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/TheLazyLemur/blitz/protocol"
	"github.com/TheLazyLemur/blitz/server/ops"
	"github.com/TheLazyLemur/blitz/types"
)

func (s *Server) handleSet(writer io.Writer, contents []byte) {
	req := &types.SetRequest{}
	err := json.Unmarshal(contents, req)
	if err != nil {
		panic(err)
	}

	ops.Set(s.Store, req.Key, req.Value)

	resp := types.SetResponse{
		Response: types.Response{
			BaseMessage: protocol.BaseMessage{
				JsonRPC: "2.0",
				Method:  "response",
			},
			ResponseCode: 1,
		},
	}

	bs, _ := protocol.EncodeMessage(resp)

	_, err = fmt.Fprintf(writer, bs)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}

func (s *Server) handleGet(writer io.Writer, contents []byte) {
	req := &types.GetRequest{}
	err := json.Unmarshal(contents, req)
	if err != nil {
		panic(err)
	}

    v, err := ops.Get(s.Store, req.Key)
    if err != nil {
        panic("fucked")
    }

	resp := types.GetResponse{
		Response: types.Response{
			BaseMessage: protocol.BaseMessage{
				JsonRPC: "2.0",
				Method:  "response",
			},
			ResponseCode: 1,
		},
        Value: v,
	}

	bs, _ := protocol.EncodeMessage(resp)

	_, err = fmt.Fprintf(writer, bs)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}
