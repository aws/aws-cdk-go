package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A load balancer that can host a VPC Endpoint Service.
type IVpcEndpointServiceLoadBalancer interface {
	// The ARN of the load balancer that hosts the VPC Endpoint Service.
	LoadBalancerArn() *string
}

// The jsii proxy for IVpcEndpointServiceLoadBalancer
type jsiiProxy_IVpcEndpointServiceLoadBalancer struct {
	_ byte // padding
}

func (j *jsiiProxy_IVpcEndpointServiceLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

