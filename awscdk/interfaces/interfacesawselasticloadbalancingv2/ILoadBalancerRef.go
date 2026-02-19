package interfacesawselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoadBalancer.
// Experimental.
type ILoadBalancerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LoadBalancer resource.
	// Experimental.
	LoadBalancerRef() *LoadBalancerReference
}

// The jsii proxy for ILoadBalancerRef
type jsiiProxy_ILoadBalancerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILoadBalancerRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ILoadBalancerRef) LoadBalancerRef() *LoadBalancerReference {
	var returns *LoadBalancerReference
	_jsii_.Get(
		j,
		"loadBalancerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoadBalancerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoadBalancerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

