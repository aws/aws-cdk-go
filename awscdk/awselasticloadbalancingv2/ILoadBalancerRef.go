package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoadBalancer.
// Experimental.
type ILoadBalancerRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILoadBalancerRef
type jsiiProxy_ILoadBalancerRef struct {
	internal.Type__constructsIConstruct
}

