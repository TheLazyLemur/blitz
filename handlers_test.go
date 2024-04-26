package main

import "testing"

func Test_parseSetCommand(t *testing.T) {
	tests := []struct {
		name        string
		raw         string
		key         string
		value       string
		timeout     int32
		expectedErr error
	}{
		{
			name:        "Miss BUTTerworth",
			raw:         "hello<|>world<|>69420",
			key:         "hello",
			value:       "world",
			timeout:     69420,
			expectedErr: nil,
		},
		{
			name:        "Miss BUTTerworth 2",
			raw:         "hello<|>world",
			key:         "",
			value:       "",
			timeout:     0,
			expectedErr: InvalidSetCommand,
		},
		{
			name:        "Miss BUTTerworth",
			raw:         "hello<|>world<|>test",
			key:         "",
			value:       "",
			timeout:     0,
			expectedErr: InvalidTimeout,
		},
		{
			name:        "Miss BUTTerworth",
			raw:         "hello<|>world<|>-10",
			key:         "",
			value:       "",
			timeout:     0,
			expectedErr: InvalidTimeout,
		},
		{
			name:        "Miss BUTTerworth",
			raw:         "hello<|>world<|>test<|>69420",
			key:         "",
			value:       "",
			timeout:     0,
			expectedErr: InvalidSetCommand,
		},
		{
			name:        "Miss BUTTerworth",
			raw:         "hello<|>world<|*|>test<|>69420",
			key:         "hello",
			value:       "world<|*|>test",
			timeout:     69420,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := parseSetCommand(tt.raw)
			if err != tt.expectedErr {
				t.Errorf("parseSetCommand() go = %v, want %v", err, tt.expectedErr)
			}

			if got != tt.key {
				t.Errorf("parseSetCommand() got = %v, want %v", got, tt.key)
			}
			if got1 != tt.value {
				t.Errorf("parseSetCommand() got1 = %v, want %v", got1, tt.value)
			}
			if got2 != tt.timeout {
				t.Errorf("parseSetCommand() got2 = %v, want %v", got2, tt.timeout)
			}
		})
	}
}
