package cf_forbidden

type Browser string

const (
	Chrome  Browser = "CHROME"
	Firefox Browser = "FIREFOX"
)

var cipherSuiteMap = map[string]uint16{
	"TLS_AES_128_GCM_SHA256":                      0x1301,
	"TLS_CHACHA20_POLY1305_SHA256":                0x1303,
	"TLS_AES_256_GCM_SHA384":                      0x1302,
	"ECDHE-ECDSA-AES128-GCM-SHA256":               0xC02B,
	"ECDHE-RSA-AES128-GCM-SHA256":                 0xC02F,
	"ECDHE-ECDSA-CHACHA20-POLY1305":               0xCCA9,
	"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256": 0xCCA8,
	"ECDHE-RSA-CHACHA20-POLY1305":                 0xCCA8,
	"ECDHE-RSA-AES256-GCM-SHA384":                 0xC030,
	"ECDHE-ECDSA-AES256-SHA":                      0xC00A,
	"ECDHE-ECDSA-AES128-SHA":                      0xC009,
	"ECDHE-RSA-AES128-SHA":                        0xC013,
	"ECDHE-RSA-AES256-SHA":                        0xC014,
	"AES128-GCM-SHA256":                           0x009C,
	"AES256-GCM-SHA384":                           0x009D,
	"AES128-SHA":                                  0x002F,
	"AES256-SHA":                                  0x0035,
}

var cipherSuites = map[Browser][]uint16{
	Chrome: {
		cipherSuiteMap["TLS_AES_128_GCM_SHA256"],
		cipherSuiteMap["TLS_CHACHA20_POLY1305_SHA256"],
		cipherSuiteMap["TLS_AES_256_GCM_SHA384"],
		cipherSuiteMap["ECDHE-ECDSA-AES128-GCM-SHA256"],
		cipherSuiteMap["ECDHE-RSA-AES128-GCM-SHA256"],
		cipherSuiteMap["ECDHE-ECDSA-CHACHA20-POLY1305"],
		cipherSuiteMap["TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256"],
		cipherSuiteMap["ECDHE-RSA-CHACHA20-POLY1305"],
		cipherSuiteMap["ECDHE-RSA-AES256-GCM-SHA384"],
		cipherSuiteMap["ECDHE-ECDSA-AES256-SHA"],
		cipherSuiteMap["ECDHE-ECDSA-AES128-SHA"],
		cipherSuiteMap["ECDHE-RSA-AES128-SHA"],
		cipherSuiteMap["ECDHE-RSA-AES256-SHA"],
		cipherSuiteMap["AES128-GCM-SHA256"],
		cipherSuiteMap["AES256-GCM-SHA384"],
		cipherSuiteMap["AES128-SHA"],
		cipherSuiteMap["AES256-SHA"],
	},
	Firefox: {
		cipherSuiteMap["TLS_AES_128_GCM_SHA256"],
		cipherSuiteMap["TLS_CHACHA20_POLY1305_SHA256"],
		cipherSuiteMap["TLS_AES_256_GCM_SHA384"],
		cipherSuiteMap["ECDHE-ECDSA-AES128-GCM-SHA256"],
		cipherSuiteMap["ECDHE-RSA-AES128-GCM-SHA256"],
		cipherSuiteMap["ECDHE-ECDSA-CHACHA20-POLY1305"],
		cipherSuiteMap["ECDHE-RSA-CHACHA20-POLY1305"],
		cipherSuiteMap["ECDHE-ECDSA-AES256-GCM-SHA384"],
		cipherSuiteMap["ECDHE-RSA-AES256-GCM-SHA384"],
		cipherSuiteMap["ECDHE-ECDSA-AES256-SHA"],
		cipherSuiteMap["ECDHE-ECDSA-AES128-SHA"],
		cipherSuiteMap["ECDHE-RSA-AES128-SHA"],
		cipherSuiteMap["ECDHE-RSA-AES256-SHA"],
		cipherSuiteMap["AES128-GCM-SHA256"],
		cipherSuiteMap["AES256-GCM-SHA384"],
		cipherSuiteMap["AES128-SHA"],
		cipherSuiteMap["AES256-SHA"],
	},
}
