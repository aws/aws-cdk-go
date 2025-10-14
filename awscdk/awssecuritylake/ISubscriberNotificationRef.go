package awssecuritylake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecuritylake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriberNotification.
// Experimental.
type ISubscriberNotificationRef interface {
	constructs.IConstruct
	// A reference to a SubscriberNotification resource.
	// Experimental.
	SubscriberNotificationRef() *SubscriberNotificationReference
}

// The jsii proxy for ISubscriberNotificationRef
type jsiiProxy_ISubscriberNotificationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISubscriberNotificationRef) SubscriberNotificationRef() *SubscriberNotificationReference {
	var returns *SubscriberNotificationReference
	_jsii_.Get(
		j,
		"subscriberNotificationRef",
		&returns,
	)
	return returns
}

