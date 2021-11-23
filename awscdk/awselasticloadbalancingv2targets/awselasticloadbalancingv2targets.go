package awselasticloadbalancingv2targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2targets/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// A single Application Load Balancer as the target for load balancing.
//
// TODO: EXAMPLE
//
// Experimental.
type AlbArnTarget interface {
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for AlbArnTarget
type jsiiProxy_AlbArnTarget struct {
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

// Create a new alb target.
// Experimental.
func NewAlbArnTarget(albArn *string, port *float64) AlbArnTarget {
	_init_.Initialize()

	j := jsiiProxy_AlbArnTarget{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.AlbArnTarget",
		[]interface{}{albArn, port},
		&j,
	)

	return &j
}

// Create a new alb target.
// Experimental.
func NewAlbArnTarget_Override(a AlbArnTarget, albArn *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.AlbArnTarget",
		[]interface{}{albArn, port},
		a,
	)
}

// Register this alb target with a load balancer.
//
// Don't call this, it is called automatically when you add the target to a
// load balancer.
// Experimental.
func (a *jsiiProxy_AlbArnTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// A single Application Load Balancer as the target for load balancing.
//
// TODO: EXAMPLE
//
// Experimental.
type AlbTarget interface {
	AlbArnTarget
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for AlbTarget
type jsiiProxy_AlbTarget struct {
	jsiiProxy_AlbArnTarget
}

// Experimental.
func NewAlbTarget(alb awselasticloadbalancingv2.ApplicationLoadBalancer, port *float64) AlbTarget {
	_init_.Initialize()

	j := jsiiProxy_AlbTarget{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.AlbTarget",
		[]interface{}{alb, port},
		&j,
	)

	return &j
}

// Experimental.
func NewAlbTarget_Override(a AlbTarget, alb awselasticloadbalancingv2.ApplicationLoadBalancer, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.AlbTarget",
		[]interface{}{alb, port},
		a,
	)
}

// Register this alb target with a load balancer.
//
// Don't call this, it is called automatically when you add the target to a
// load balancer.
// Experimental.
func (a *jsiiProxy_AlbTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// An EC2 instance that is the target for load balancing.
//
// If you register a target of this type, you are responsible for making
// sure the load balancer's security group can connect to the instance.
//
// TODO: EXAMPLE
//
// Experimental.
type InstanceIdTarget interface {
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for InstanceIdTarget
type jsiiProxy_InstanceIdTarget struct {
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

// Create a new Instance target.
// Experimental.
func NewInstanceIdTarget(instanceId *string, port *float64) InstanceIdTarget {
	_init_.Initialize()

	j := jsiiProxy_InstanceIdTarget{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.InstanceIdTarget",
		[]interface{}{instanceId, port},
		&j,
	)

	return &j
}

// Create a new Instance target.
// Experimental.
func NewInstanceIdTarget_Override(i InstanceIdTarget, instanceId *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.InstanceIdTarget",
		[]interface{}{instanceId, port},
		i,
	)
}

// Register this instance target with a load balancer.
//
// Don't call this, it is called automatically when you add the target to a
// load balancer.
// Experimental.
func (i *jsiiProxy_InstanceIdTarget) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Register this instance target with a load balancer.
//
// Don't call this, it is called automatically when you add the target to a
// load balancer.
// Experimental.
func (i *jsiiProxy_InstanceIdTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
// Experimental.
type InstanceTarget interface {
	InstanceIdTarget
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for InstanceTarget
type jsiiProxy_InstanceTarget struct {
	jsiiProxy_InstanceIdTarget
}

// Create a new Instance target.
// Experimental.
func NewInstanceTarget(instance awsec2.Instance, port *float64) InstanceTarget {
	_init_.Initialize()

	j := jsiiProxy_InstanceTarget{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.InstanceTarget",
		[]interface{}{instance, port},
		&j,
	)

	return &j
}

// Create a new Instance target.
// Experimental.
func NewInstanceTarget_Override(i InstanceTarget, instance awsec2.Instance, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.InstanceTarget",
		[]interface{}{instance, port},
		i,
	)
}

// Register this instance target with a load balancer.
//
// Don't call this, it is called automatically when you add the target to a
// load balancer.
// Experimental.
func (i *jsiiProxy_InstanceTarget) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Register this instance target with a load balancer.
//
// Don't call this, it is called automatically when you add the target to a
// load balancer.
// Experimental.
func (i *jsiiProxy_InstanceTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

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
// TODO: EXAMPLE
//
// Experimental.
type IpTarget interface {
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for IpTarget
type jsiiProxy_IpTarget struct {
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
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
// Experimental.
func NewIpTarget(ipAddress *string, port *float64, availabilityZone *string) IpTarget {
	_init_.Initialize()

	j := jsiiProxy_IpTarget{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.IpTarget",
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
// Experimental.
func NewIpTarget_Override(i IpTarget, ipAddress *string, port *float64, availabilityZone *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.IpTarget",
		[]interface{}{ipAddress, port, availabilityZone},
		i,
	)
}

// Register this instance target with a load balancer.
//
// Don't call this, it is called automatically when you add the target to a
// load balancer.
// Experimental.
func (i *jsiiProxy_IpTarget) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Register this instance target with a load balancer.
//
// Don't call this, it is called automatically when you add the target to a
// load balancer.
// Experimental.
func (i *jsiiProxy_IpTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
// Experimental.
type LambdaTarget interface {
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for LambdaTarget
type jsiiProxy_LambdaTarget struct {
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
}

// Create a new Lambda target.
// Experimental.
func NewLambdaTarget(fn awslambda.IFunction) LambdaTarget {
	_init_.Initialize()

	j := jsiiProxy_LambdaTarget{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.LambdaTarget",
		[]interface{}{fn},
		&j,
	)

	return &j
}

// Create a new Lambda target.
// Experimental.
func NewLambdaTarget_Override(l LambdaTarget, fn awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_targets.LambdaTarget",
		[]interface{}{fn},
		l,
	)
}

// Register this instance target with a load balancer.
//
// Don't call this, it is called automatically when you add the target to a
// load balancer.
// Experimental.
func (l *jsiiProxy_LambdaTarget) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		l,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Register this instance target with a load balancer.
//
// Don't call this, it is called automatically when you add the target to a
// load balancer.
// Experimental.
func (l *jsiiProxy_LambdaTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		l,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

