package protocol

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"log/slog"
)

func ParseRequest(r io.Reader) (Command, error) {
	var keyLen int32
	if err := binary.Read(r, binary.LittleEndian, &keyLen); err != nil {
		if err == io.EOF {
			slog.Info("Connection closed by client")
		}
		return Command{}, err
	}

	cmdBytes := make([]byte, keyLen)
	if err := binary.Read(r, binary.LittleEndian, &cmdBytes); err != nil {
		return Command{}, err
	}

	request := &Command{}
	if err := json.Unmarshal(cmdBytes, request); err != nil {
		slog.Error("failed to unmarshal json", "error", err.Error())
		return Command{}, err
	}

	return *request, nil
}


func SendCommand(w io.Writer, cmd Command) error {
	buf := new(bytes.Buffer)

	bytes, _ := json.Marshal(cmd)
	str := string(bytes)

	keyLen := int32(len(str))
	if err := binary.Write(buf, binary.LittleEndian, keyLen); err != nil {
		return err
	}

	if err := binary.Write(buf, binary.LittleEndian, []byte(str)); err != nil {
		return err
	}

	_, err := w.Write(buf.Bytes())
	return err
}

func ReadResponse(r io.Reader) (Response, error) {
	var kLen int32
	if err := binary.Read(r, binary.LittleEndian, &kLen); err != nil {
		return Response{}, err
	}

	cmdBytes := make([]byte, kLen)
	if err := binary.Read(r, binary.LittleEndian, &cmdBytes); err != nil {
		return Response{}, err
	}

	result := new(Response)
	_ = json.Unmarshal(cmdBytes, result)
	return *result, nil
}

func SendResponse(w io.Writer, result Response) {
	resultJson, _ := json.Marshal(result)
	resLen := len(resultJson)

	if err := binary.Write(w, binary.LittleEndian, int32(resLen)); err != nil {
		if err == io.EOF {
			slog.Info("Connection closed by client")
		}
	}

	if err := binary.Write(w, binary.LittleEndian, resultJson); err != nil {
		if err == io.EOF {
			slog.Info("Connection closed by client")
		}
	}
}
