package awsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventType.
// Experimental.
type IEventTypeRef interface {
	constructs.IConstruct
	// A reference to a EventType resource.
	// Experimental.
	EventTypeRef() *EventTypeReference
}

// The jsii proxy for IEventTypeRef
type jsiiProxy_IEventTypeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEventTypeRef) EventTypeRef() *EventTypeReference {
	var returns *EventTypeReference
	_jsii_.Get(
		j,
		"eventTypeRef",
		&returns,
	)
	return returns
}

