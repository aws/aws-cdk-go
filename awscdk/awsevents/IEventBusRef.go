package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventBus.
// Experimental.
type IEventBusRef interface {
	constructs.IConstruct
	// A reference to a EventBus resource.
	// Experimental.
	EventBusRef() *EventBusReference
}

// The jsii proxy for IEventBusRef
type jsiiProxy_IEventBusRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEventBusRef) EventBusRef() *EventBusReference {
	var returns *EventBusReference
	_jsii_.Get(
		j,
		"eventBusRef",
		&returns,
	)
	return returns
}

