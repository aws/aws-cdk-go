package awscustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventStream.
// Experimental.
type IEventStreamRef interface {
	constructs.IConstruct
	// A reference to a EventStream resource.
	// Experimental.
	EventStreamRef() *EventStreamReference
}

// The jsii proxy for IEventStreamRef
type jsiiProxy_IEventStreamRef struct {
	internal.Type__constructsIConstruct
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

