package util

import (
	"encoding/base64"
)

// https://viewblock.io/arweave/tokens
var PopularPSTAddress = map[string]string{
	"ARDRIVE": "-8A6RexFkpfWwuyVO98wzSFZh0d6VJuI-buTJvlwOJQ",
	"VRT":     "usjm4PCxUd5mtaon7zc97-dt-3qf67yPyqgzLnLqk5A",
}

func IsValidAddress(input string) bool {
	// An arweave address base64url encode of a SHA-256 hash)
	bytes, err := base64.RawURLEncoding.DecodeString(input)
	if err != nil {
		return false
	}
	// Valid base64url, check if bytes length must be 32 (SHA-256 hash)
	if len(bytes) != 32 {
		return false
	} else {
		return true
	}
}
