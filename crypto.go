package verifier

import (
	"crypto/sha256"
	"math/big"
)

type ecdsaSignature struct {
	R, S *big.Int
}

func sha256Hash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

