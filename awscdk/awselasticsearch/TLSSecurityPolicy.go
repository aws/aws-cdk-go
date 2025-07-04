package awselasticsearch


// The minimum TLS version required for traffic to the domain.
// Deprecated: use opensearchservice module instead.
type TLSSecurityPolicy string

const (
	// Cipher suite TLS 1.0.
	// Deprecated: use opensearchservice module instead.
	TLSSecurityPolicy_TLS_1_0 TLSSecurityPolicy = "TLS_1_0"
	// Cipher suite TLS 1.2.
	// Deprecated: use opensearchservice module instead.
	TLSSecurityPolicy_TLS_1_2 TLSSecurityPolicy = "TLS_1_2"
)

