package cxapi


// Load balancer ip address type.
type LoadBalancerIpAddressType string

const (
	// IPV4 ip address.
	LoadBalancerIpAddressType_IPV4 LoadBalancerIpAddressType = "IPV4"
	// Dual stack address.
	LoadBalancerIpAddressType_DUAL_STACK LoadBalancerIpAddressType = "DUAL_STACK"
)

