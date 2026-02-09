package cxapi


// Load balancer ip address type.
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type LoadBalancerIpAddressType string

const (
	// IPV4 ip address.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	LoadBalancerIpAddressType_IPV4 LoadBalancerIpAddressType = "IPV4"
	// Dual stack address.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	LoadBalancerIpAddressType_DUAL_STACK LoadBalancerIpAddressType = "DUAL_STACK"
	// IPv6 only public addresses, with private IPv4 and IPv6 addresses.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	LoadBalancerIpAddressType_DUAL_STACK_WITHOUT_PUBLIC_IPV4 LoadBalancerIpAddressType = "DUAL_STACK_WITHOUT_PUBLIC_IPV4"
)

