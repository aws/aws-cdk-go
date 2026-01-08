package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Indicates that this resource can be referenced as an ALB Listener.
type IApplicationListenerRef interface {
	IListener
	// Indicates that this is an ALB listener.
	//
	// Will always return true, but is necessary to prevent accidental structural
	// equality in TypeScript.
	IsApplicationListener() *bool
}

// The jsii proxy for IApplicationListenerRef
type jsiiProxy_IApplicationListenerRef struct {
	jsiiProxy_IListener
}

func (j *jsiiProxy_IApplicationListenerRef) IsApplicationListener() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isApplicationListener",
		&returns,
	)
	return returns
}

