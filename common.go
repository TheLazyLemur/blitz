package main

import "io"

// handleError handles errors returned by various functions and
// determines the appropriate action to take.
func handleError(err error) error {
	if err == io.EOF {
		return ErrConnectionClosedByClient
	}

	return err
}
