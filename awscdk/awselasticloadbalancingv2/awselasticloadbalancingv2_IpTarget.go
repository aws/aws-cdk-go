package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An IP address that is a target for load balancing.
//
// Specify IP addresses from the subnets of the virtual private cloud (VPC) for
// the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and
// 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify
// publicly routable IP addresses.
//
// If you register a target of this type, you are responsible for making
// sure the load balancer's security group can send packets to the IP address.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipTarget := awscdk.Aws_elasticloadbalancingv2.NewIpTarget(jsii.String("ipAddress"), jsii.Number(123), jsii.String("availabilityZone"))
//
// Deprecated: Use IpTarget from the.
type IpTarget interface {
	IApplicationLoadBalancerTarget
	INetworkLoadBalancerTarget
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	// Deprecated: Use IpTarget from the.
	AttachToApplicationTargetGroup(targetGroup IApplicationTargetGroup) *LoadBalancerTargetProps
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	// Deprecated: Use IpTarget from the.
	AttachToNetworkTargetGroup(targetGroup INetworkTargetGroup) *LoadBalancerTargetProps
}

// The jsii proxy struct for IpTarget
type jsiiProxy_IpTarget struct {
	jsiiProxy_IApplicationLoadBalancerTarget
	jsiiProxy_INetworkLoadBalancerTarget
}

// Create a new IPAddress target.
//
// The availabilityZone parameter determines whether the target receives
// traffic from the load balancer nodes in the specified Availability Zone
// or from all enabled Availability Zones for the load balancer.
//
// This parameter is not supported if the target type of the target group
// is instance. If the IP address is in a subnet of the VPC for the target
// group, the Availability Zone is automatically detected and this
// parameter is optional. If the IP address is outside the VPC, this
// parameter is required.
//
// With an Application Load Balancer, if the IP address is outside the VPC
// for the target group, the only supported value is all.
//
// Default is automatic.
// Deprecated: Use IpTarget from the.
func NewIpTarget(ipAddress *string, port *float64, availabilityZone *string) IpTarget {
	_init_.Initialize()

	if err := validateNewIpTargetParameters(ipAddress); err != nil {
		panic(err)
	}
	j := jsiiProxy_IpTarget{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2.IpTarget",
		[]interface{}{ipAddress, port, availabilityZone},
		&j,
	)

	return &j
}

// Create a new IPAddress target.
//
// The availabilityZone parameter determines whether the target receives
// traffic from the load balancer nodes in the specified Availability Zone
// or from all enabled Availability Zones for the load balancer.
//
// This parameter is not supported if the target type of the target group
// is instance. If the IP address is in a subnet of the VPC for the target
// group, the Availability Zone is automatically detected and this
// parameter is optional. If the IP address is outside the VPC, this
// parameter is required.
//
// With an Application Load Balancer, if the IP address is outside the VPC
// for the target group, the only supported value is all.
//
// Default is automatic.
// Deprecated: Use IpTarget from the.
func NewIpTarget_Override(i IpTarget, ipAddress *string, port *float64, availabilityZone *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2.IpTarget",
		[]interface{}{ipAddress, port, availabilityZone},
		i,
	)
}

func (i *jsiiProxy_IpTarget) AttachToApplicationTargetGroup(targetGroup IApplicationTargetGroup) *LoadBalancerTargetProps {
	if err := i.validateAttachToApplicationTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpTarget) AttachToNetworkTargetGroup(targetGroup INetworkTargetGroup) *LoadBalancerTargetProps {
	if err := i.validateAttachToNetworkTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

