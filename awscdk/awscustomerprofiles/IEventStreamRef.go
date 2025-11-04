package awscustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventStream.
// Experimental.
type IEventStreamRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EventStream resource.
	// Experimental.
	EventStreamRef() *EventStreamReference
}

// The jsii proxy for IEventStreamRef
type jsiiProxy_IEventStreamRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEventStreamRef) EventStreamRef() *EventStreamReference {
	var returns *EventStreamReference
	_jsii_.Get(
		j,
		"eventStreamRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventStreamRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventStreamRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

