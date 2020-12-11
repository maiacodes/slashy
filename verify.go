package main

import (
	"crypto/ed25519"
	"encoding/hex"
)

func verify(signature, hash, publicKey string) bool {
	decodedSig, err := hex.DecodeString(signature)
	if err != nil {
		panic(err)
	}

	decodedPubKey, err := hex.DecodeString(publicKey)
	if err != nil {
		panic(err)
	}

	return ed25519.Verify(decodedPubKey, []byte(hash), decodedSig)
}
