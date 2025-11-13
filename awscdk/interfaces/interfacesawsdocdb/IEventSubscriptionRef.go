package interfacesawsdocdb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdocdb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventSubscription.
// Experimental.
type IEventSubscriptionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EventSubscription resource.
	// Experimental.
	EventSubscriptionRef() *EventSubscriptionReference
}

// The jsii proxy for IEventSubscriptionRef
type jsiiProxy_IEventSubscriptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IEventSubscriptionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventSubscriptionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

