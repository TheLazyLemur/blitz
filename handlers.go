package main

import (
	"encoding/binary"
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	InvalidSetCommand = errors.New("Invalid set command")
	InvalidTimeout    = errors.New("Invalid timeout format")
)

type Handlers interface {
	HandleSet(io.Reader, io.Writer) error
	HandleGet(io.Reader, io.Writer) error
}

type handlersImpl struct{}

// parseSetCommand function KEY<|>VALUE<|>Timeout
func parseSetCommand(raw string) (string, string, int32, error) {
	parts := strings.Split(raw, "<|>")
	if len(parts) != 3 {
		return "", "", 0, InvalidSetCommand
	}

	timeout, err := strconv.Atoi(parts[2])
	if err != nil || timeout <= 0 {
		return "", "", 0, InvalidTimeout
	}

	return parts[0], parts[1], int32(timeout), nil
}

func (h *handlersImpl) HandleSet(r io.Reader, w io.Writer) error {
	var msgLen int32
	if err := binary.Read(r, binary.LittleEndian, &msgLen); err != nil {
		return handleError(err)
	}

	raw := make([]byte, msgLen)
	_ = binary.Read(r, binary.LittleEndian, &raw)

	key, value, timeout, err := parseSetCommand(string(raw))
	if err != nil {
		return handleError(err)
	}

	_ = key
	_ = value
	_ = timeout

	if err := binary.Write(w, binary.LittleEndian, Ok); err != nil {
		return handleError(err)
	}

	return nil
}

func (h *handlersImpl) HandleGet(_ io.Reader, w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, Ok); err != nil {
		return handleError(err)
	}

	return nil
}
