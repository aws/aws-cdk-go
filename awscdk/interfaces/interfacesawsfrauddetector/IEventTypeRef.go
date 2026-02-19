package interfacesawsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventType.
// Experimental.
type IEventTypeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EventType resource.
	// Experimental.
	EventTypeRef() *EventTypeReference
}

// The jsii proxy for IEventTypeRef
type jsiiProxy_IEventTypeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEventTypeRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEventTypeRef) EventTypeRef() *EventTypeReference {
	var returns *EventTypeReference
	_jsii_.Get(
		j,
		"eventTypeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventTypeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventTypeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

