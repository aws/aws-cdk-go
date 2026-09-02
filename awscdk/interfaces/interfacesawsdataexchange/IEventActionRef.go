package interfacesawsdataexchange

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdataexchange/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventAction.
// Experimental.
type IEventActionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EventAction resource.
	// Experimental.
	EventActionRef() *EventActionReference
}

// The jsii proxy for IEventActionRef
type jsiiProxy_IEventActionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEventActionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEventActionRef) EventActionRef() *EventActionReference {
	var returns *EventActionReference
	_jsii_.Get(
		j,
		"eventActionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventActionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventActionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

