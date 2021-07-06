package verifiertest

import "strings"

// PrepareByteArray is a helper function to create a delimiter-joined byte array from given string array
func PrepareByteArray(data []string) []byte {
	return []byte(strings.Join(data, "\u2063"))
}
