package util

import (
	"testing"
)

func TestARToWinston(t *testing.T) {
	tests := []struct {
		arAmount string
		output   string
	}{
		{
			arAmount: "1",
			output:   "1000000000000",
		},
		{
			arAmount: "200",
			output:   "200000000000000",
		},
		{
			arAmount: "1.0",
			output:   "1000000000000",
		},
		{
			arAmount: "1.33",
			output:   "1330000000000",
		},
		{
			arAmount: "1.123456789012",
			output:   "1123456789012",
		},
		{
			arAmount: "1.",
			output:   "1000000000000",
		},
		{
			arAmount: ".3",
			output:   "300000000000",
		},
	}

	for i, tc := range tests {
		got, _ := ARToWinston(tc.arAmount)
		if tc.output != got {
			t.Fatalf("test case %d: expected: %v, got: %v", i+1, tc.output, got)
		}
	}
}
