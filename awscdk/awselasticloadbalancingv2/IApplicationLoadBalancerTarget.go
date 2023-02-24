package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for constructs that can be targets of an application load balancer.
type IApplicationLoadBalancerTarget interface {
	// Attach load-balanced target to a TargetGroup.
	//
	// May return JSON to directly add to the [Targets] list, or return undefined
	// if the target will register itself with the load balancer.
	AttachToApplicationTargetGroup(targetGroup IApplicationTargetGroup) *LoadBalancerTargetProps
}

// The jsii proxy for IApplicationLoadBalancerTarget
type jsiiProxy_IApplicationLoadBalancerTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IApplicationLoadBalancerTarget) AttachToApplicationTargetGroup(targetGroup IApplicationTargetGroup) *LoadBalancerTargetProps {
	if err := i.validateAttachToApplicationTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

