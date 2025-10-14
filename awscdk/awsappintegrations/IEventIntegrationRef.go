package awsappintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappintegrations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventIntegration.
// Experimental.
type IEventIntegrationRef interface {
	constructs.IConstruct
	// A reference to a EventIntegration resource.
	// Experimental.
	EventIntegrationRef() *EventIntegrationReference
}

// The jsii proxy for IEventIntegrationRef
type jsiiProxy_IEventIntegrationRef struct {
	internal.Type__constructsIConstruct
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

