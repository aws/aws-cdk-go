package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A network target group.
type INetworkTargetGroup interface {
	INetworkTargetGroupRef
	ITargetGroup
	// Add a load balancing target to this target group.
	AddTarget(targets ...INetworkLoadBalancerTarget)
	// Register a listener that is load balancing to this target group.
	//
	// Don't call this directly. It will be called by listeners.
	RegisterListener(listener INetworkListenerRef)
	// All metrics available for this target group.
	Metrics() INetworkTargetGroupMetrics
}

// The jsii proxy for INetworkTargetGroup
type jsiiProxy_INetworkTargetGroup struct {
	jsiiProxy_INetworkTargetGroupRef
	jsiiProxy_ITargetGroup
}

func (i *jsiiProxy_INetworkTargetGroup) AddTarget(targets ...INetworkLoadBalancerTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addTarget",
		args,
	)
}

func (i *jsiiProxy_INetworkTargetGroup) RegisterListener(listener INetworkListenerRef) {
	if err := i.validateRegisterListenerParameters(listener); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerListener",
		[]interface{}{listener},
	)
}

func (i *jsiiProxy_INetworkTargetGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_INetworkTargetGroup) Metrics() INetworkTargetGroupMetrics {
	var returns INetworkTargetGroupMetrics
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkTargetGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkTargetGroup) IsNetworkTargetGroup() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isNetworkTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkTargetGroup) LoadBalancerArns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkTargetGroup) LoadBalancerAttached() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"loadBalancerAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkTargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkTargetGroup) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkTargetGroup) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkTargetGroup) TargetGroupRef() *interfacesawselasticloadbalancingv2.TargetGroupReference {
	var returns *interfacesawselasticloadbalancingv2.TargetGroupReference
	_jsii_.Get(
		j,
		"targetGroupRef",
		&returns,
	)
	return returns
}

