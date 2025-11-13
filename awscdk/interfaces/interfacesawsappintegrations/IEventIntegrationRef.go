package interfacesawsappintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappintegrations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventIntegration.
// Experimental.
type IEventIntegrationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EventIntegration resource.
	// Experimental.
	EventIntegrationRef() *EventIntegrationReference
}

// The jsii proxy for IEventIntegrationRef
type jsiiProxy_IEventIntegrationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IEventIntegrationRef) EventIntegrationRef() *EventIntegrationReference {
	var returns *EventIntegrationReference
	_jsii_.Get(
		j,
		"eventIntegrationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventIntegrationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventIntegrationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

