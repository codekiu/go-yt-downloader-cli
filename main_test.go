package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCliFlags(t *testing.T) {
	oldArgs := os.Args
	oldFlagCommandLine := flag.CommandLine
	defer func() {
		os.Args = oldArgs
		flag.CommandLine = oldFlagCommandLine
	}()

	tests := []struct {
		expectedErr error
		expected    Config
		name        string
		args        []string
	}{
		{
			name: "Should work",
			args: []string{"cmd", "-f", "mp3", "-u", "test"},
			expected: Config{
				ytUrl:  "test",
				format: "mp3",
			},
			expectedErr: nil,
		},
		{
			name: "Wrong format",
			args: []string{"cmd", "-f", "mp5", "-u", "test"},
			expected: Config{
				ytUrl:  "test",
				format: "mp3",
			},
			expectedErr: fmt.Errorf("unsupported format: mp5"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock os args
			os.Args = tt.args
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

			config, err := parseInput()

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, config)
			}
		})
	}
}
