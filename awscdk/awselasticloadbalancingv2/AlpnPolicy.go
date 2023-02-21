package awselasticloadbalancingv2


// Application-Layer Protocol Negotiation Policies for network load balancers.
//
// Which protocols should be used over a secure connection.
type AlpnPolicy string

const (
	// Negotiate only HTTP/1.*. The ALPN preference list is http/1.1, http/1.0.
	AlpnPolicy_HTTP1_ONLY AlpnPolicy = "HTTP1_ONLY"
	// Negotiate only HTTP/2.
	//
	// The ALPN preference list is h2.
	AlpnPolicy_HTTP2_ONLY AlpnPolicy = "HTTP2_ONLY"
	// Prefer HTTP/1.* over HTTP/2 (which can be useful for HTTP/2 testing). The ALPN preference list is http/1.1, http/1.0, h2.
	AlpnPolicy_HTTP2_OPTIONAL AlpnPolicy = "HTTP2_OPTIONAL"
	// Prefer HTTP/2 over HTTP/1.*. The ALPN preference list is h2, http/1.1, http/1.0.
	AlpnPolicy_HTTP2_PREFERRED AlpnPolicy = "HTTP2_PREFERRED"
	// Do not negotiate ALPN.
	AlpnPolicy_NONE AlpnPolicy = "NONE"
)

