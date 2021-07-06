package verifier

import (
	"crypto/ecdsa"
	"encoding/asn1"
	"encoding/base64"
	"errors"
	"strings"

	certs "github.com/google/certificate-transparency-go"
)

const (
	// ApplePublicKey is Apple's public key for SkAdNetwork postback verification.
	ApplePublicKey = "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEWdp8GPcGqmhgzEFj9Z2nSpQVddayaPe4FMzqM9wib1+aHaaIzoHoLN9zW4K8y4SPykE3YVK3sVqW6Af0lfx3gg=="

	// AppleDelimiter is the delimiter used for joining parameters
	AppleDelimiter = "\u2063"
)

var (
	// ErrorInvalidData indicates input data is invalid
	ErrorInvalidData       = errors.New("invalid data")

	// ErrorSignatureMismatch indicates signature did not match input data
	ErrorSignatureMismatch = errors.New("signature mismatch")
)

// Verifier verifies that a given postback data matches its signature
type Verifier struct {
	key *ecdsa.PublicKey
}

// NewVerifier is used to create an instance of Verifier for postback versions 2.1 and later.
func NewVerifier() *Verifier {
	var key, _ = certs.PublicKeyFromB64(ApplePublicKey)
	return &Verifier{key.(*ecdsa.PublicKey)}
}

// VerifySignature verifies a given postback data and signature
// The parameter data is expected to be a byte array. If string array is given, the function will use the correct
// delimiter to create the byte array needed to verify the signature against.
func (v *Verifier) VerifySignature(data interface{}, signature string) error {
	var b []byte

	switch v := data.(type) {
	case []string:
		b = []byte(strings.Join(v, AppleDelimiter))
	case []byte:
		b = v
	default:
		return ErrorInvalidData
	}

	return v.verifySignature(b, signature)
}

func (v *Verifier) verifySignature(data []byte, base64der string) error {
	der, err := base64.StdEncoding.DecodeString(base64der)
	if err != nil {
		return err
	}

	signature := &ecdsaSignature{}
	if _, err = asn1.Unmarshal(der, signature); err != nil {
		return err
	}

	if valid := ecdsa.Verify(v.key, sha256Hash(data), signature.R, signature.S); valid {
		return nil
	}

	return ErrorSignatureMismatch
}
