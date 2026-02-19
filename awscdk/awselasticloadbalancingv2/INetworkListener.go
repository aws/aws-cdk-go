package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties to reference an existing listener.
type INetworkListener interface {
	IListener
	INetworkListenerRef
}

// The jsii proxy for INetworkListener
type jsiiProxy_INetworkListener struct {
	jsiiProxy_IListener
	jsiiProxy_INetworkListenerRef
}

func (i *jsiiProxy_INetworkListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_INetworkListener) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INetworkListener) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkListener) IsNetworkListener() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isNetworkListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkListener) ListenerRef() *interfacesawselasticloadbalancingv2.ListenerReference {
	var returns *interfacesawselasticloadbalancingv2.ListenerReference
	_jsii_.Get(
		j,
		"listenerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

