package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for listener actions.
type IListenerAction interface {
	// Render the actions in this chain.
	RenderActions() *[]*CfnListener_ActionProperty
}

// The jsii proxy for IListenerAction
type jsiiProxy_IListenerAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IListenerAction) RenderActions() *[]*CfnListener_ActionProperty {
	var returns *[]*CfnListener_ActionProperty

	_jsii_.Invoke(
		i,
		"renderActions",
		nil, // no parameters
		&returns,
	)

	return returns
}

