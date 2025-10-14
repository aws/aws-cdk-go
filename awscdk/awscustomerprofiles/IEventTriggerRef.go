package awscustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventTrigger.
// Experimental.
type IEventTriggerRef interface {
	constructs.IConstruct
	// A reference to a EventTrigger resource.
	// Experimental.
	EventTriggerRef() *EventTriggerReference
}

// The jsii proxy for IEventTriggerRef
type jsiiProxy_IEventTriggerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEventTriggerRef) EventTriggerRef() *EventTriggerReference {
	var returns *EventTriggerReference
	_jsii_.Get(
		j,
		"eventTriggerRef",
		&returns,
	)
	return returns
}

