package awscustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventTrigger.
// Experimental.
type IEventTriggerRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EventTrigger resource.
	// Experimental.
	EventTriggerRef() *EventTriggerReference
}

// The jsii proxy for IEventTriggerRef
type jsiiProxy_IEventTriggerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IEventTriggerRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventTriggerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

