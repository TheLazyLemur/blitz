package main

import "errors"

var (
	ErrConnectionClosedByClient = errors.New("Client closed connection")
	ErrInvalidCommand           = errors.New("Invalid command")
)

type Command uint8

const (
	VoidCommand Command = iota
	Set
	Get
)

func (c Command) IsValid() bool {
	switch c {
	case VoidCommand:
		return false
	case Set:
		return true
	case Get:
		return true
	default:
		return false
	}
}

type Status uint8

const (
	VoidStatus Status = iota
	Ok
	InvalidCommand
)
