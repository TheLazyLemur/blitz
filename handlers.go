package main

import (
	"encoding/binary"
	"io"
)

func handleSet(_ io.Reader, w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, Ok); err != nil {
		return handleError(ErrInvalidCommand)
	}

	return nil
}

func handleGet(_ io.Reader, w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, Ok); err != nil {
		return handleError(ErrInvalidCommand)
	}

	return nil
}
