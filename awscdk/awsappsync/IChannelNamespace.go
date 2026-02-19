package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AppSync channel namespace.
type IChannelNamespace interface {
	interfacesawsappsync.IChannelNamespaceRef
	awscdk.IResource
	// The ARN of the AppSync channel namespace.
	ChannelNamespaceArn() *string
}

// The jsii proxy for IChannelNamespace
type jsiiProxy_IChannelNamespace struct {
	internal.Type__interfacesawsappsyncIChannelNamespaceRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IChannelNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IChannelNamespace) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IChannelNamespace) ChannelNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelNamespace) ChannelNamespaceRef() *interfacesawsappsync.ChannelNamespaceReference {
	var returns *interfacesawsappsync.ChannelNamespaceReference
	_jsii_.Get(
		j,
		"channelNamespaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelNamespace) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

