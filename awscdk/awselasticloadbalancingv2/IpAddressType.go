package awselasticloadbalancingv2


// What kind of addresses to allocate to the load balancer.
//
// Example:
//   var vpc vpc
//
//
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("LB"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
//   	EnablePrefixForIpv6SourceNat: jsii.Boolean(true),
//   })
//
//   listener := lb.AddListener(jsii.String("Listener"), &BaseNetworkListenerProps{
//   	Port: jsii.Number(1229),
//   	Protocol: elbv2.Protocol_UDP,
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

