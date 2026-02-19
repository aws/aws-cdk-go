package awselasticloadbalancing

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancing"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a load balancer.
type ILoadBalancer interface {
	interfacesawselasticloadbalancing.ILoadBalancerRef
	awscdk.IResource
}

// The jsii proxy for ILoadBalancer
type jsiiProxy_ILoadBalancer struct {
	internal.Type__interfacesawselasticloadbalancingILoadBalancerRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ILoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ILoadBalancer) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ILoadBalancer) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoadBalancer) LoadBalancerRef() *interfacesawselasticloadbalancing.LoadBalancerReference {
	var returns *interfacesawselasticloadbalancing.LoadBalancerReference
	_jsii_.Get(
		j,
		"loadBalancerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

