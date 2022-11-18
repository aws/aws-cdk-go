package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2/internal"
)

// A target group.
// Experimental.
type ITargetGroup interface {
	awscdk.IConstruct
	// A token representing a list of ARNs of the load balancers that route traffic to this target group.
	// Experimental.
	LoadBalancerArns() *string
	// Return an object to depend on the listeners added to this target group.
	// Experimental.
	LoadBalancerAttached() awscdk.IDependable
	// ARN of the target group.
	// Experimental.
	TargetGroupArn() *string
	// The name of the target group.
	// Experimental.
	TargetGroupName() *string
}

// The jsii proxy for ITargetGroup
type jsiiProxy_ITargetGroup struct {
	internal.Type__awscdkIConstruct
}

func (j *jsiiProxy_ITargetGroup) LoadBalancerArns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetGroup) LoadBalancerAttached() awscdk.IDependable {
	var returns awscdk.IDependable
	_jsii_.Get(
		j,
		"loadBalancerAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetGroup) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetGroup) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

