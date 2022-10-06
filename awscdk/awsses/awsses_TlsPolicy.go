package awsses


// The type of TLS policy for a receipt rule.
// Experimental.
type TlsPolicy string

const (
	// Do not check for TLS.
	// Experimental.
	TlsPolicy_OPTIONAL TlsPolicy = "OPTIONAL"
	// Bounce emails that are not received over TLS.
	// Experimental.
	TlsPolicy_REQUIRE TlsPolicy = "REQUIRE"
)

