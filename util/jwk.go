package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/everFinance/gojwk"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

func GenerateRsaJwk() (jwk.Key, error) {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, fmt.Errorf(`failed to generate RSA private key: %w`, err)
	}

	k, err := jwk.FromRaw(key)
	if err != nil {
		return nil, fmt.Errorf(`failed to generate jwk.RSAPrivateKey: %w`, err)
	}

	return k, nil
}

func JwkToAddress(b []byte) (string, error) {
	key, err := gojwk.Unmarshal(b)
	if err != nil {
		return "", fmt.Errorf(`failed to unmarshal jwk: %w`, err)
	}

	pubKey, err := key.DecodePublicKey()
	if err != nil {
		return "", fmt.Errorf(`failed to decode public key: %w`, err)
	}
	pub, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("pubKey type error")
	}
	// The Arweave address is the SHA-256 hash of the RSA public key.
	addr := sha256.Sum256(pub.N.Bytes())
	return base64.RawURLEncoding.EncodeToString(addr[:]), nil
}
