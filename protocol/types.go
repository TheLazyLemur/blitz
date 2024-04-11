package protocol

var (
	Separator = []byte{'\r', '\n', '\r', '\n'}
)

type BaseMessage struct {
	Method  string `json:"method"`
	JsonRPC string `json:"jsonrpc"`
}
