package awselasticloadbalancingv2targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// A single Application Load Balancer as the target for load balancing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationLoadBalancerRef IApplicationLoadBalancerRef
//
//   albTarget := awscdk.Aws_elasticloadbalancingv2_targets.NewAlbTarget(applicationLoadBalancerRef, jsii.Number(123))
//
// Deprecated: Use `AlbListenerTarget` instead or
// `AlbArnTarget` for an imported load balancer. This target does not automatically
// add a dependency between the ALB listener and resulting NLB target group,
// without which may cause stack deployments to fail if the NLB target group is provisioned
// before the listener has been fully created.
type AlbTarget interface {
	AlbArnTarget
	// Register this alb target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	// Deprecated: Use `AlbListenerTarget` instead or
	// `AlbArnTarget` for an imported load balancer. This target does not automatically
	// add a dependency between the ALB listener and resulting NLB target group,
	// without which may cause stack deployments to fail if the NLB target group is provisioned
	// before the listener has been fully created.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for AlbTarget
type jsiiProxy_AlbTarget struct {
	jsiiProxy_AlbArnTarget
}

// Deprecated: Use `AlbListenerTarget` instead or
// `AlbArnTarget` for an imported load balancer. This target does not automatically
// add a dependency between the ALB listener and resulting NLB target group,
// without which may cause stack deployments to fail if the NLB target group is provisioned
// before the listener has been fully created.
func NewAlbTarget(alb awselasticloadbalancingv2.IApplicationLoadBalancerRef, port *float64) AlbTarget {
	_init_.Initialize()

	if err := validateNewAlbTargetParameters(alb, port); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.AlbTarget",
		[]interface{}{alb, port},
		&j,
	)

	return &j
}

// Deprecated: Use `AlbListenerTarget` instead or
// `AlbArnTarget` for an imported load balancer. This target does not automatically
// add a dependency between the ALB listener and resulting NLB target group,
// without which may cause stack deployments to fail if the NLB target group is provisioned
// before the listener has been fully created.
func NewAlbTarget_Override(a AlbTarget, alb awselasticloadbalancingv2.IApplicationLoadBalancerRef, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.AlbTarget",
		[]interface{}{alb, port},
		a,
	)
}

func (a *jsiiProxy_AlbTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := a.validateAttachToNetworkTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

