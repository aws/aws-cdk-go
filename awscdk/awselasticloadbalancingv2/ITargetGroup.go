package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A target group.
type ITargetGroup interface {
	constructs.IConstruct
	interfacesawselasticloadbalancingv2.ITargetGroupRef
	// A token representing a list of ARNs of the load balancers that route traffic to this target group.
	LoadBalancerArns() *string
	// Return an object to depend on the listeners added to this target group.
	LoadBalancerAttached() constructs.IDependable
	// ARN of the target group.
	TargetGroupArn() *string
	// The name of the target group.
	TargetGroupName() *string
}

// The jsii proxy for ITargetGroup
type jsiiProxy_ITargetGroup struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesawselasticloadbalancingv2ITargetGroupRef
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

func (j *jsiiProxy_ITargetGroup) LoadBalancerAttached() constructs.IDependable {
	var returns constructs.IDependable
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

func (j *jsiiProxy_ITargetGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetGroup) TargetGroupRef() *interfacesawselasticloadbalancingv2.TargetGroupReference {
	var returns *interfacesawselasticloadbalancingv2.TargetGroupReference
	_jsii_.Get(
		j,
		"targetGroupRef",
		&returns,
	)
	return returns
}

