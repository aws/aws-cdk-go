package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A network target group.
type INetworkTargetGroup interface {
	ITargetGroup
	// Add a load balancing target to this target group.
	AddTarget(targets ...INetworkLoadBalancerTarget)
	// Register a listener that is load balancing to this target group.
	//
	// Don't call this directly. It will be called by listeners.
	RegisterListener(listener INetworkListener)
}

// The jsii proxy for INetworkTargetGroup
type jsiiProxy_INetworkTargetGroup struct {
	jsiiProxy_ITargetGroup
}

func (i *jsiiProxy_INetworkTargetGroup) AddTarget(targets ...INetworkLoadBalancerTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addTarget",
		args,
	)
}

func (i *jsiiProxy_INetworkTargetGroup) RegisterListener(listener INetworkListener) {
	if err := i.validateRegisterListenerParameters(listener); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerListener",
		[]interface{}{listener},
	)
}

