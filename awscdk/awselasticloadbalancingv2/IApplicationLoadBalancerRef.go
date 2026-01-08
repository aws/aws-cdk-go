package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2"
)

// Indicates that this resource can be referenced as an Application LoadBalancer.
type IApplicationLoadBalancerRef interface {
	interfacesawselasticloadbalancingv2.ILoadBalancerRef
	// Indicates that this is an Application Load Balancer.
	//
	// Will always return true, but is necessary to prevent accidental structural
	// equality in TypeScript.
	IsApplicationLoadBalancer() *bool
}

// The jsii proxy for IApplicationLoadBalancerRef
type jsiiProxy_IApplicationLoadBalancerRef struct {
	internal.Type__interfacesawselasticloadbalancingv2ILoadBalancerRef
}

func (j *jsiiProxy_IApplicationLoadBalancerRef) IsApplicationLoadBalancer() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isApplicationLoadBalancer",
		&returns,
	)
	return returns
}

