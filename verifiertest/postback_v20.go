package verifiertest

import "fmt"

// data from https://developer.apple.com/documentation/storekit/skadnetwork/verifying_an_install-validation_postback
var data20 = map[string]interface{}{
	"version":               "2.0",
	"ad-network-id":         "com.example",
	"campaign-id":           42,
	"transaction-id":        "6aafb7a5-0170-41b5-bbe4-fe71dedf1e28",
	"app-id":                525463029,
	"attribution-signature": "MDYCGQD0AdGn5gUnSuVGk8Wi0IgxzWiKdBzwJrQCGQCJfkrI5bda93EC4Xm1H+MtNxstFmnVBn0=",
	"redownload":            true,
	"source-app-id":         1234567891,
	"timestamp":             "12342352346",
	"conversion-value":      20,
}

// PostbackV20 returns valid data and signature for v2.0 postback
func PostbackV20() ([]string, string) {
	return []string{
			fmt.Sprintf("%v", data20["version"]),
			fmt.Sprintf("%v", data20["ad-network-id"]),
			fmt.Sprintf("%v", data20["campaign-id"]),
			fmt.Sprintf("%v", data20["app-id"]),
			fmt.Sprintf("%v", data20["transaction-id"]),
			fmt.Sprintf("%v", data20["redownload"]),
			fmt.Sprintf("%v", data20["source-app-id"]),
		},
		fmt.Sprintf("%v", data20["attribution-signature"])
}
