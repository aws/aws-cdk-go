package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Interface for ECS load balancer target.
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

