package awsdocdb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdocdb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventSubscription.
// Experimental.
type IEventSubscriptionRef interface {
	constructs.IConstruct
	// A reference to a EventSubscription resource.
	// Experimental.
	EventSubscriptionRef() *EventSubscriptionReference
}

// The jsii proxy for IEventSubscriptionRef
type jsiiProxy_IEventSubscriptionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEventSubscriptionRef) EventSubscriptionRef() *EventSubscriptionReference {
	var returns *EventSubscriptionReference
	_jsii_.Get(
		j,
		"eventSubscriptionRef",
		&returns,
	)
	return returns
}

