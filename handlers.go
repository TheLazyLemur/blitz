package main

import (
	"encoding/binary"
	"io"
)

type Handlers interface {
	HandleSet(io.Reader, io.Writer) error
	HandleGet(io.Reader, io.Writer) error
}

type handlersImpl struct{}

func (h *handlersImpl) HandleSet(_ io.Reader, w io.Writer) error {
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
