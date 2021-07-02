package verifier

import (
	"github.com/tenjin/go-sk-postback-verifier/verifiertest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyPostback(t *testing.T) {
	v := NewVerifier()

	data30, signature30 := verifiertest.PostbackV30()
	data20, signature20 := verifiertest.PostbackV20()

	// With v3.0 postback data
	err := v.VerifySignature(data30, signature30)
	assert.NoError(t, err, "Sample data from Apple should be valid")

	// Same but with byte array
	b30 := verifiertest.PrepareByteArray(data30)
	err = v.VerifySignature(b30, signature30)
	assert.NoError(t, err, "Sample data from Apple should be valid")

	// With legacy postback data
	err = v.VerifySignature(data20, signature20)
	assert.Error(t, err, "Legacy postback sample from Apple should not be valid with current verifier")

	// Same but with byte array
	b20 := verifiertest.PrepareByteArray(data20)
	err = v.VerifySignature(b20, signature20)
	assert.Error(t, err, "Legacy postback sample from Apple should not be valid with current verifier")
}

func TestVerifyLegacyPostback(t *testing.T) {
	v := NewLegacyVerifier()

	data30, signature30 := verifiertest.PostbackV30()
	data20, signature20 := verifiertest.PostbackV20()

	// With v3.0 postback data
	err := v.VerifySignature(data30, signature30)
	assert.Error(t, err, "Recent sample data should not be verifiable with the legacy verifier")

	// Same but with byte array
	b30 := verifiertest.PrepareByteArray(data30)
	err = v.VerifySignature(b30, signature30)
	assert.Error(t, err, "Recent sample data should not be verifiable with the legacy verifier")

	// With legacy postback data
	err = v.VerifySignature(data20, signature20)
	assert.NoError(t, err, "Legacy postback sample from Apple should be valid with the legacy verifier")

	// Same but with byte array
	b20 := verifiertest.PrepareByteArray(data20)
	err = v.VerifySignature(b20, signature20)
	assert.NoError(t, err, "Legacy postback sample from Apple should be valid with the legacy verifier")
}

func TestVerifyInvalidPostback(t *testing.T) {
	v := NewVerifier()

	// Given bad data string array
	data, signature := verifiertest.PostbackV30()
	data = append(data, "invalid data")

	// ErrorSignatureMismatch is returned
	err := v.VerifySignature(data, signature)
	assert.ErrorIs(t, err, ErrorSignatureMismatch, "Invalid data")

	// Given bad data format
	dataInt := []int{1, 2, 3}

	// ErrorInvalidData is returned
	err = v.VerifySignature(dataInt, signature)
	assert.ErrorIs(t, err, ErrorInvalidData, "Invalid data")

	// Given invalid signature
	data, signature = verifiertest.PostbackV30()
	signature += "x"

	// Error is returned
	err = v.VerifySignature(data, signature)
	assert.Error(t, err, "Invalid signature")
}
