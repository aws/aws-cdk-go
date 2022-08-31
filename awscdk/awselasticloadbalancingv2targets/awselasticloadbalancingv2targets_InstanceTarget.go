package awselasticloadbalancingv2targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instance instance
//
//   instanceTarget := awscdk.Aws_elasticloadbalancingv2_targets.NewInstanceTarget(instance, jsii.Number(123))
//
// Experimental.
type InstanceTarget interface {
	InstanceIdTarget
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	// Experimental.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	// Experimental.
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

