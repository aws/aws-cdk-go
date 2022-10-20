package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for classes that provide the peer-specification parts of a security group rule.
type IPeer interface {
	IConnectable
	// Produce the egress rule JSON for the given connection.
	ToEgressRuleConfig() interface{}
	// Produce the ingress rule JSON for the given connection.
	ToIngressRuleConfig() interface{}
	// Whether the rule can be inlined into a SecurityGroup or not.
	CanInlineRule() *bool
	// A unique identifier for this connection peer.
	UniqueId() *string
}

// The jsii proxy for IPeer
type jsiiProxy_IPeer struct {
	jsiiProxy_IConnectable
}

func (i *jsiiProxy_IPeer) ToEgressRuleConfig() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toEgressRuleConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPeer) ToIngressRuleConfig() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toIngressRuleConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPeer) CanInlineRule() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canInlineRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPeer) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}

