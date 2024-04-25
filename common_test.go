package main

import (
	"errors"
	"io"
	"testing"
)

func Test_handleError(t *testing.T) {
	testErr := errors.New("Test error")

	tests := []struct {
		name        string
		err         error
		expectedErr error
	}{
		{
			name:        "Transformed Error - io.EOF",
			err:         io.EOF,
			expectedErr: ErrConnectionClosedByClient,
		},
		{
			name:        "Non Transformed Error - io.EOF",
			err:         testErr,
			expectedErr: testErr,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := handleError(tc.err); err != tc.expectedErr {
				t.Errorf("handleError() error = %v, wantErr %v", err, tc.expectedErr)
			}
		})
	}
}
