package interfacesawsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventBus.
// Experimental.
type IEventBusRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EventBus resource.
	// Experimental.
	EventBusRef() *EventBusReference
}

// The jsii proxy for IEventBusRef
type jsiiProxy_IEventBusRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEventBusRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IEventBusRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBusRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

