package verifier

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
)

const (
	// ApplePublicKeyLegacy has been deprecated by Apple, but can be used for verifying postbacks versions 1.0 and 2.0.
	ApplePublicKeyLegacy = "MEkwEwYHKoZIzj0CAQYIKoZIzj0DAQEDMgAEMyHD625uvsmGq4C43cQ9BnfN2xslVT5V1nOmAMP6qaRRUll3PB1JYmgSm+62sosG"
)

// NewLegacyVerifier can be used for postbacks versions 1.0 and 2.0. Google's certificate-transparency-go does not work
// with ApplePublicKeyLegacy but we can create it manually.
func NewLegacyVerifier() *Verifier {
	// The following corresponds to running `certs.PublicKeyFromB64(ApplePublicKeyLegacy)`
	x, _ := new(big.Int).SetString("1253750435644680222299188317470676239292798058580822802005", 10)
	y, _ := new(big.Int).SetString("5258341652717200205739849880119579488055200286266702859014", 10)
	p, _ := new(big.Int).SetString("6277101735386680763835789423207666416083908700390324961279", 10)
	n, _ := new(big.Int).SetString("6277101735386680763835789423176059013767194773182842284081", 10)
	b, _ := new(big.Int).SetString("2455155546008943817740293915197451784769108058161191238065", 10)
	gx, _ := new(big.Int).SetString("602046282375688656758213480587526111916698976636884684818", 10)
	gy, _ := new(big.Int).SetString("174050332293622031404857552280219410364023488927386650641", 10)
	key := &ecdsa.PublicKey{
		Curve: &elliptic.CurveParams{P: p, N: n, B: b, Gx: gx, Gy: gy, BitSize: 192, Name: "P-192"},
		X:     x,
		Y:     y,
	}
	return &Verifier{key}
}
