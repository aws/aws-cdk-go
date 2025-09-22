package awselasticloadbalancingv2


// The IP address type of targets registered with a target group.
//
// Example:
//   var vpc vpc
//
//
//   ipv4ApplicationTargetGroup := elbv2.NewApplicationTargetGroup(this, jsii.String("IPv4ApplicationTargetGroup"), &ApplicationTargetGroupProps{
//   	Vpc: Vpc,
//   	Port: jsii.Number(80),
//   	TargetType: elbv2.TargetType_INSTANCE,
//   	IpAddressType: elbv2.TargetGroupIpAddressType_IPV4,
//   })
//
//   ipv6ApplicationTargetGroup := elbv2.NewApplicationTargetGroup(this, jsii.String("Ipv6ApplicationTargetGroup"), &ApplicationTargetGroupProps{
//   	Vpc: Vpc,
//   	Port: jsii.Number(80),
//   	TargetType: elbv2.TargetType_INSTANCE,
//   	IpAddressType: elbv2.TargetGroupIpAddressType_IPV6,
//   })
//
//   ipv4NetworkTargetGroup := elbv2.NewNetworkTargetGroup(this, jsii.String("IPv4NetworkTargetGroup"), &NetworkTargetGroupProps{
//   	Vpc: Vpc,
//   	Port: jsii.Number(80),
//   	TargetType: elbv2.TargetType_INSTANCE,
//   	IpAddressType: elbv2.TargetGroupIpAddressType_IPV4,
//   })
//
//   ipv6NetworkTargetGroup := elbv2.NewNetworkTargetGroup(this, jsii.String("Ipv6NetworkTargetGroup"), &NetworkTargetGroupProps{
//   	Vpc: Vpc,
//   	Port: jsii.Number(80),
//   	TargetType: elbv2.TargetType_INSTANCE,
//   	IpAddressType: elbv2.TargetGroupIpAddressType_IPV6,
//   })
//
type TargetGroupIpAddressType string

const (
	// IPv4 addresses.
	TargetGroupIpAddressType_IPV4 TargetGroupIpAddressType = "IPV4"
	// IPv6 addresses.
	TargetGroupIpAddressType_IPV6 TargetGroupIpAddressType = "IPV6"
)

