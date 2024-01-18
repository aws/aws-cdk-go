package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for adding new network targets to a listener.
//
// Example:
//   var vpc vpc
//   var asg autoScalingGroup
//   var sg1 iSecurityGroup
//   var sg2 iSecurityGroup
//
//
//   // Create the load balancer in a VPC. 'internetFacing' is 'false'
//   // by default, which creates an internal load balancer.
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("LB"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   	SecurityGroups: []*iSecurityGroup{
//   		sg1,
//   	},
//   })
//   lb.AddSecurityGroup(sg2)
//
//   // Add a listener on a particular port.
//   listener := lb.AddListener(jsii.String("Listener"), &BaseNetworkListenerProps{
//   	Port: jsii.Number(443),
//   })
//
//   // Add targets on a particular port.
//   listener.AddTargets(jsii.String("AppFleet"), &AddNetworkTargetsProps{
//   	Port: jsii.Number(443),
//   	Targets: []iNetworkLoadBalancerTarget{
//   		asg,
//   	},
//   })
//
type AddNetworkTargetsProps struct {
	// The port on which the listener listens for requests.
	// Default: Determined from protocol if known.
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	// Default: Duration.minutes(5)
	//
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	// Default: - The default value for each property in this configuration varies depending on the target.
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Indicates whether client IP preservation is enabled.
	// Default: false if the target group type is IP address and the
	// target group protocol is TCP or TLS. Otherwise, true.
	//
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Protocol for target group, expects TCP, TLS, UDP, or TCP_UDP.
	// Default: - inherits the protocol of the listener.
	//
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Indicates whether Proxy Protocol version 2 is enabled.
	// Default: false.
	//
	ProxyProtocolV2 *bool `field:"optional" json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	// Default: Automatically generated.
	//
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	Targets *[]INetworkLoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

