package cxapi


// Load balancer ip address type.
// Experimental.
type LoadBalancerIpAddressType string

const (
	// IPV4 ip address.
	// Experimental.
	LoadBalancerIpAddressType_IPV4 LoadBalancerIpAddressType = "IPV4"
	// Dual stack address.
	// Experimental.
	LoadBalancerIpAddressType_DUAL_STACK LoadBalancerIpAddressType = "DUAL_STACK"
)

