package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "normal",
			input:  "hiepdinh",
			output: "hnidpeih",
		},
		{
			name:   "unicode",
			input:  "hiepdinh了。",
			output: "。了hnidpeih",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if reverseString := reverseString(v.input); reverseString != v.output {
				t.Errorf("Revert %v expect %v but got %v", v.input, v.output, reverseString)
			}
		})
	}
}
