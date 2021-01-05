package main

import (
	"crypto/ed25519"
	"encoding/hex"
)

func verify(signature, hash, publicKey string) bool {
	decodedSig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	decodedPubKey, err := hex.DecodeString(publicKey)
	if err != nil {
		return false
	}

	return ed25519.Verify(decodedPubKey, []byte(hash), decodedSig)
}
