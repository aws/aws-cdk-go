package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Indicates that this resource can be referenced as an NLB Listener.
type INetworkListenerRef interface {
	IListener
	// Indicates that this is an NLB listener.
	//
	// Will always return true, but is necessary to prevent accidental structural
	// equality in TypeScript.
	IsNetworkListener() *bool
}

// The jsii proxy for INetworkListenerRef
type jsiiProxy_INetworkListenerRef struct {
	jsiiProxy_IListener
}

func (j *jsiiProxy_INetworkListenerRef) IsNetworkListener() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isNetworkListener",
		&returns,
	)
	return returns
}

