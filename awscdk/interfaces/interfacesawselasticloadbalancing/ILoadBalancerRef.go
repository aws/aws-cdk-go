package interfacesawselasticloadbalancing

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancing/internal"
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

