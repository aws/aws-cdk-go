package interfacesawsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventBusPolicy.
// Experimental.
type IEventBusPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EventBusPolicy resource.
	// Experimental.
	EventBusPolicyRef() *EventBusPolicyReference
}

// The jsii proxy for IEventBusPolicyRef
type jsiiProxy_IEventBusPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEventBusPolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEventBusPolicyRef) EventBusPolicyRef() *EventBusPolicyReference {
	var returns *EventBusPolicyReference
	_jsii_.Get(
		j,
		"eventBusPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBusPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBusPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

