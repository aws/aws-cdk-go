package awselasticloadbalancing

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing/internal"
)

// Interface that is going to be implemented by constructs that you can load balance to.
type ILoadBalancerTarget interface {
	awsec2.IConnectable
	// Attach load-balanced target to a classic ELB.
	AttachToClassicLB(loadBalancer LoadBalancer)
}

// The jsii proxy for ILoadBalancerTarget
type jsiiProxy_ILoadBalancerTarget struct {
	internal.Type__awsec2IConnectable
}

func (i *jsiiProxy_ILoadBalancerTarget) AttachToClassicLB(loadBalancer LoadBalancer) {
	if err := i.validateAttachToClassicLBParameters(loadBalancer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

