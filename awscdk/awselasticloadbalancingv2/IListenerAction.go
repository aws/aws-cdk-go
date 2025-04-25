package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for listener actions.
type IListenerAction interface {
	// Render the listener default actions in this chain.
	RenderActions() *[]*CfnListener_ActionProperty
	// Render the listener rule actions in this chain.
	RenderRuleActions() *[]*CfnListenerRule_ActionProperty
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

func (i *jsiiProxy_IListenerAction) RenderRuleActions() *[]*CfnListenerRule_ActionProperty {
	var returns *[]*CfnListenerRule_ActionProperty

	_jsii_.Invoke(
		i,
		"renderRuleActions",
		nil, // no parameters
		&returns,
	)

	return returns
}

