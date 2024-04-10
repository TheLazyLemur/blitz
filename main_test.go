package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"testing"
)

type EOFReader struct{}

func (r *EOFReader) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func sampleClientInput(c Command) *bytes.Reader {
	inputBuf := &bytes.Buffer{}
	binary.Write(inputBuf, binary.LittleEndian, c)
	return bytes.NewReader(inputBuf.Bytes())
}

func Test_handleNetworkCommunication(t *testing.T) {
	tcs := []struct {
		name             string
		command          Command
		inputFunc        func(c Command) *bytes.Reader
		expectedError    error
		expectedResponse Status
		EOF              bool
	}{
		{
			name:             "Success_Set",
			command:          Set,
			inputFunc:        sampleClientInput,
			expectedError:    nil,
			expectedResponse: Ok,
		},
		{
			name:             "Success_Get",
			command:          Get,
			inputFunc:        sampleClientInput,
			expectedError:    nil,
			expectedResponse: Ok,
		},
		{
			name:             "Fail_Invalid_Command",
			command:          Command(123),
			inputFunc:        sampleClientInput,
			expectedError:    ErrInvalidCommand,
			expectedResponse: InvalidCommand,
		},
		{
			name:             "Fail_EOF_Err",
			command:          Command(123),
			expectedError:    ErrConnectionClosedByClient,
			expectedResponse: InvalidCommand,
			EOF:              true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if tc.EOF {
				err := handleNetworkCommunication(&EOFReader{}, nil)
				if err != tc.expectedError {
					t.Error("Expected", tc.expectedError, "Got", err)
				}
				return
			}

			inputReader := tc.inputFunc(tc.command)
			outputBuf := &bytes.Buffer{}
			err := handleNetworkCommunication(inputReader, outputBuf)
			if err != tc.expectedError {
				t.Error("Expected", tc.expectedError, "Got", err)
			}

			var replyFromServer Status
			_ = binary.Read(outputBuf, binary.LittleEndian, &replyFromServer)

			if replyFromServer != tc.expectedResponse {
				t.Error("Expected", tc.expectedResponse, "Got", replyFromServer)
			}
		})
	}
}

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
