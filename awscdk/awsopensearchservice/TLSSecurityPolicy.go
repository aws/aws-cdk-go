package awsopensearchservice


// The minimum TLS version required for traffic to the domain.
type TLSSecurityPolicy string

const (
	// Cipher suite TLS 1.0.
	TLSSecurityPolicy_TLS_1_0 TLSSecurityPolicy = "TLS_1_0"
	// Cipher suite TLS 1.2.
	TLSSecurityPolicy_TLS_1_2 TLSSecurityPolicy = "TLS_1_2"
)

