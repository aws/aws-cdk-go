package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventInvokeConfig.
// Experimental.
type IEventInvokeConfigRef interface {
	constructs.IConstruct
	// A reference to a EventInvokeConfig resource.
	// Experimental.
	EventInvokeConfigRef() *EventInvokeConfigReference
}

// The jsii proxy for IEventInvokeConfigRef
type jsiiProxy_IEventInvokeConfigRef struct {
	internal.Type__constructsIConstruct
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

