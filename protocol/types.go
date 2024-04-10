package protocol

type Command struct {
	Command string
	Key     string
	Value   string
}

func NewCommand(c, k, v string) Command {
	return Command{
		Command: c,
		Key:     k,
		Value:   v,
	}
}

type Response struct {
	StatusCode int32
	Value      string
}

func NewResult(statusCode int32, value []byte) Response {
	return Response{
		StatusCode: statusCode,
		Value:      string(value),
	}
}
