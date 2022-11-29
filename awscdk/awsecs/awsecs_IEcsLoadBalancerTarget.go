package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

// Interface for ECS load balancer target.
// Experimental.
type IEcsLoadBalancerTarget interface {
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancing.ILoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
}

// The jsii proxy for IEcsLoadBalancerTarget
type jsiiProxy_IEcsLoadBalancerTarget struct {
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
	internal.Type__awselasticloadbalancingILoadBalancerTarget
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

