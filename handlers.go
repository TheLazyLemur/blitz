package main

import (
	"encoding/binary"
	"io"
)

type handlers struct{}

func (h *handlers) handleSet(_ io.Reader, w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, Ok); err != nil {
		return handleError(err)
	}

	return nil
}

func (h *handlers) handleGet(_ io.Reader, w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, Ok); err != nil {
		return handleError(err)
	}

	return nil
}
