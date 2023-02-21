package awselasticloadbalancingv2


// What kind of addresses to allocate to the load balancer.
// Experimental.
type IpAddressType string

const (
	// Allocate IPv4 addresses.
	// Experimental.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// Allocate both IPv4 and IPv6 addresses.
	// Experimental.
	IpAddressType_DUAL_STACK IpAddressType = "DUAL_STACK"
)

