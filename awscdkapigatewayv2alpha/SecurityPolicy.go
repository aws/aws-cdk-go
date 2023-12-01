package awscdkapigatewayv2alpha


// The minimum version of the SSL protocol that you want API Gateway to use for HTTPS connections.
// Deprecated.
type SecurityPolicy string

const (
	// Cipher suite TLS 1.0.
	// Deprecated.
	SecurityPolicy_TLS_1_0 SecurityPolicy = "TLS_1_0"
	// Cipher suite TLS 1.2.
	// Deprecated.
	SecurityPolicy_TLS_1_2 SecurityPolicy = "TLS_1_2"
)

