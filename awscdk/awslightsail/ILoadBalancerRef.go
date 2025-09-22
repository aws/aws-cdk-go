package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoadBalancer.
// Experimental.
type ILoadBalancerRef interface {
	constructs.IConstruct
	// A reference to a LoadBalancer resource.
	// Experimental.
	LoadBalancerRef() *LoadBalancerReference
}

// The jsii proxy for ILoadBalancerRef
type jsiiProxy_ILoadBalancerRef struct {
	internal.Type__constructsIConstruct
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

