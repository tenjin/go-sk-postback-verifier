package verifiertest

import (
	"fmt"
)

// data from https://developer.apple.com/documentation/storekit/skadnetwork/verifying_an_install-validation_postback
var data30 = map[string]interface{}{
	"version":               "3.0",
	"ad-network-id":         "com.example",
	"campaign-id":           42,
	"transaction-id":        "f9ac267a-a889-44ce-b5f7-0166d11461f0",
	"app-id":                525463029,
	"attribution-signature": "MEUCIQDYfConaAkeeGvAr6WAjBbY7LBX1z6ir/8T4jVYKJaMPQIgHC5jsV0lvlaWgFr7ON0VN4rmTTW9gZUzFoLkEn/g+g8=",
	"redownload":            true,
	"fidelity-type":         1,
	"did-win":               false,
}

// PostbackV30 returns valid data and signature for v3.0 postback
func PostbackV30() ([]string, string) {
	return []string{
			fmt.Sprintf("%v", data30["version"]),
			fmt.Sprintf("%v", data30["ad-network-id"]),
			fmt.Sprintf("%v", data30["campaign-id"]),
			fmt.Sprintf("%v", data30["app-id"]),
			fmt.Sprintf("%v", data30["transaction-id"]),
			fmt.Sprintf("%v", data30["redownload"]),
			fmt.Sprintf("%v", data30["fidelity-type"]),
			fmt.Sprintf("%v", data30["did-win"]),
		},
		fmt.Sprintf("%v", data30["attribution-signature"])
}
