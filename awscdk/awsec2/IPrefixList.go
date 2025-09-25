package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A prefix list.
type IPrefixList interface {
	IPeer
	IPrefixListRef
	awscdk.IResource
	// The ID of the prefix list.
	PrefixListId() *string
}

// The jsii proxy for IPrefixList
type jsiiProxy_IPrefixList struct {
	jsiiProxy_IPeer
	jsiiProxy_IPrefixListRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPrefixList) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IPrefixList) ToEgressRuleConfig() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toEgressRuleConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPrefixList) ToIngressRuleConfig() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toIngressRuleConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPrefixList) PrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixList) CanInlineRule() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canInlineRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixList) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixList) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixList) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixList) PrefixListRef() *PrefixListReference {
	var returns *PrefixListReference
	_jsii_.Get(
		j,
		"prefixListRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixList) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixList) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}

