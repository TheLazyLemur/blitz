package ops

import (
	"fmt"

	"github.com/TheLazyLemur/blitz/protocol"
)

var (
    db = make(map[string][]byte)
)

func HandleSet(command protocol.Command) {
    db[command.Key] = []byte(command.Value)
    fmt.Println(len(db))
}

func HandleGet(command protocol.Command) []byte {
    return db[command.Key]
}
