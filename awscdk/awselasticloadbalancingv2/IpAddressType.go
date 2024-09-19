package awselasticloadbalancingv2


// What kind of addresses to allocate to the load balancer.
//
// Example:
//   var vpc vpc
//
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
//   })
//
type IpAddressType string

const (
	// Allocate IPv4 addresses.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// Allocate both IPv4 and IPv6 addresses.
	IpAddressType_DUAL_STACK IpAddressType = "DUAL_STACK"
	// IPv6 only public addresses, with private IPv4 and IPv6 addresses.
	IpAddressType_DUAL_STACK_WITHOUT_PUBLIC_IPV4 IpAddressType = "DUAL_STACK_WITHOUT_PUBLIC_IPV4"
)

