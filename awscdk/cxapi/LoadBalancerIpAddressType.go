package cxapi


// Load balancer ip address type.
type LoadBalancerIpAddressType string

const (
	// IPV4 ip address.
	LoadBalancerIpAddressType_IPV4 LoadBalancerIpAddressType = "IPV4"
	// Dual stack address.
	LoadBalancerIpAddressType_DUAL_STACK LoadBalancerIpAddressType = "DUAL_STACK"
	// IPv6 only public addresses, with private IPv4 and IPv6 addresses.
	LoadBalancerIpAddressType_DUAL_STACK_WITHOUT_PUBLIC_IPV4 LoadBalancerIpAddressType = "DUAL_STACK_WITHOUT_PUBLIC_IPV4"
)

