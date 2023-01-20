package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for constructs that can be targets of an network load balancer.
type INetworkLoadBalancerTarget interface {
	// Attach load-balanced target to a TargetGroup.
	//
	// May return JSON to directly add to the [Targets] list, or return undefined
	// if the target will register itself with the load balancer.
	AttachToNetworkTargetGroup(targetGroup INetworkTargetGroup) *LoadBalancerTargetProps
}

// The jsii proxy for INetworkLoadBalancerTarget
type jsiiProxy_INetworkLoadBalancerTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_INetworkLoadBalancerTarget) AttachToNetworkTargetGroup(targetGroup INetworkTargetGroup) *LoadBalancerTargetProps {
	if err := i.validateAttachToNetworkTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

