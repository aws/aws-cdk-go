package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
)

type ILoadBalancerV2 interface {
	awscdk.IResource
	// The canonical hosted zone ID of this load balancer.
	//
	// Example value: `Z2P70J7EXAMPLE`.
	LoadBalancerCanonicalHostedZoneId() *string
	// The DNS name of this load balancer.
	//
	// Example value: `my-load-balancer-424835706.us-west-2.elb.amazonaws.com`
	LoadBalancerDnsName() *string
}

// The jsii proxy for ILoadBalancerV2
type jsiiProxy_ILoadBalancerV2 struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILoadBalancerV2) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoadBalancerV2) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

