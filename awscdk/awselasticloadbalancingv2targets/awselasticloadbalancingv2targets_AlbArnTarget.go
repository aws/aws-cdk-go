package awselasticloadbalancingv2targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2targets/internal"
)

// A single Application Load Balancer as the target for load balancing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   albArnTarget := awscdk.Aws_elasticloadbalancingv2_targets.NewAlbArnTarget(jsii.String("albArn"), jsii.Number(123))
//
// Experimental.
type AlbArnTarget interface {
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	// Register this alb target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	// Experimental.
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

	if err := validateNewAlbArnTargetParameters(albArn, port); err != nil {
		panic(err)
	}
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

func (a *jsiiProxy_AlbArnTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
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

