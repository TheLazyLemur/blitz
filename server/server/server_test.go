package server

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"testing"

	"github.com/TheLazyLemur/blitz/protocol"
)

func Test_handleCommand(t *testing.T) {
	tcs := []struct {
		name     string
		cmd      protocol.Command
		expected protocol.Response
	}{
		{
			name: "Test set command",
			cmd: protocol.Command{
				Command: "set",
				Key:     "key",
				Value:   "value",
			},
			expected: protocol.Response{
				StatusCode: 0,
				Value:      "",
			},
		},
		{
			name: "Test get command",
			cmd: protocol.Command{
				Command: "get",
				Key:     "key",
			},
			expected: protocol.Response{
				StatusCode: 0,
				Value:      "value",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)

            s := NewServer(Config{})
			s.handleCommand(buf, tc.cmd)

			var kLen int32
			if err := binary.Read(buf, binary.LittleEndian, &kLen); err != nil {
                t.Fail()
			}

			cmdBytes := make([]byte, kLen)
			if err := binary.Read(buf, binary.LittleEndian, &cmdBytes); err != nil {
                t.Fail()
			}

			actual := new(protocol.Response)
			_ = json.Unmarshal(cmdBytes, actual)

			if *actual != tc.expected {
				t.Fatal()
			}
		})
	}

}
