package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglobalaccelerator"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface of the Listener.
type IListener interface {
	interfacesawsglobalaccelerator.IListenerRef
	awscdk.IResource
	// The ARN of the listener.
	ListenerArn() *string
}

// The jsii proxy for IListener
type jsiiProxy_IListener struct {
	internal.Type__interfacesawsglobalacceleratorIListenerRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IListener) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IListener) ListenerRef() *interfacesawsglobalaccelerator.ListenerReference {
	var returns *interfacesawsglobalaccelerator.ListenerReference
	_jsii_.Get(
		j,
		"listenerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

