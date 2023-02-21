package awsses


// The type of TLS policy for a receipt rule.
type TlsPolicy string

const (
	// Do not check for TLS.
	TlsPolicy_OPTIONAL TlsPolicy = "OPTIONAL"
	// Bounce emails that are not received over TLS.
	TlsPolicy_REQUIRE TlsPolicy = "REQUIRE"
)

