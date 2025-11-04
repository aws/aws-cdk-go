package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventInvokeConfig.
// Experimental.
type IEventInvokeConfigRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EventInvokeConfig resource.
	// Experimental.
	EventInvokeConfigRef() *EventInvokeConfigReference
}

// The jsii proxy for IEventInvokeConfigRef
type jsiiProxy_IEventInvokeConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEventInvokeConfigRef) EventInvokeConfigRef() *EventInvokeConfigReference {
	var returns *EventInvokeConfigReference
	_jsii_.Get(
		j,
		"eventInvokeConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventInvokeConfigRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventInvokeConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

