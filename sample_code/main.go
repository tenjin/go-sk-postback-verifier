package main

import (
	"fmt"
	"github.com/tenjin/go-sk-postback-verifier"
)

// go run main.go
func main() {
	v := verifier.NewVerifier()

	// parameters to be verified in order
	p := []string{
		"3.0",
		"com.example",
		"42",
		"525463029",
		"f9ac267a-a889-44ce-b5f7-0166d11461f0",
		"true",
		"1",
		"false",
	}

	// "attribution-signature" field
	s := "MEUCIQDYfConaAkeeGvAr6WAjBbY7LBX1z6ir/8T4jVYKJaMPQIgHC5jsV0lvlaWgFr7ON0VN4rmTTW9gZUzFoLkEn/g+g8="

	err := v.VerifySignature(p, s)
	if err != nil {
		fmt.Println("FAIL")
	} else {
		fmt.Println("PASS")
	}
}
