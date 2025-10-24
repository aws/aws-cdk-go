package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for a new Network Target Group.
//
// Example:
//   var vpc Vpc
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
type NetworkTargetGroupProps struct {
	// Indicates whether cross zone load balancing is enabled.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattribute.html
	//
	// Default: - use load balancer configuration.
	//
	CrossZoneEnabled *bool `field:"optional" json:"crossZoneEnabled" yaml:"crossZoneEnabled"`
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	// Default: 300.
	//
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	// Default: - The default value for each property in this configuration varies depending on the target.
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The type of IP addresses of the targets registered with the target group.
	// Default: undefined - ELB defaults to IPv4.
	//
	IpAddressType TargetGroupIpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Configuring target group health.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
	//
	// Default: - use default configuration.
	//
	TargetGroupHealth *TargetGroupHealth `field:"optional" json:"targetGroupHealth" yaml:"targetGroupHealth"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	// Default: - Automatically generated.
	//
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	// Default: - Determined automatically.
	//
	TargetType TargetType `field:"optional" json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`.
	// Default: - undefined.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The port on which the target receives traffic.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Indicates whether the load balancer terminates connections at the end of the deregistration timeout.
	// Default: false.
	//
	ConnectionTermination *bool `field:"optional" json:"connectionTermination" yaml:"connectionTermination"`
	// Indicates whether client IP preservation is enabled.
	// Default: false if the target group type is IP address and the
	// target group protocol is TCP or TLS. Otherwise, true.
	//
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Protocol for target group, expects TCP, TLS, UDP, or TCP_UDP.
	// Default: - TCP.
	//
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Indicates whether Proxy Protocol version 2 is enabled.
	// Default: false.
	//
	ProxyProtocolV2 *bool `field:"optional" json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	// Default: - No targets.
	//
	Targets *[]INetworkLoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

