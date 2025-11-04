package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventSourceMapping.
// Experimental.
type IEventSourceMappingRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EventSourceMapping resource.
	// Experimental.
	EventSourceMappingRef() *EventSourceMappingReference
}

// The jsii proxy for IEventSourceMappingRef
type jsiiProxy_IEventSourceMappingRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEventSourceMappingRef) EventSourceMappingRef() *EventSourceMappingReference {
	var returns *EventSourceMappingReference
	_jsii_.Get(
		j,
		"eventSourceMappingRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventSourceMappingRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventSourceMappingRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

