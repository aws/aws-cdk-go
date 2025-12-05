package interfacesawsobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TelemetryPipelines.
// Experimental.
type ITelemetryPipelinesRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TelemetryPipelines resource.
	// Experimental.
	TelemetryPipelinesRef() *TelemetryPipelinesReference
}

// The jsii proxy for ITelemetryPipelinesRef
type jsiiProxy_ITelemetryPipelinesRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITelemetryPipelinesRef) TelemetryPipelinesRef() *TelemetryPipelinesReference {
	var returns *TelemetryPipelinesReference
	_jsii_.Get(
		j,
		"telemetryPipelinesRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITelemetryPipelinesRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITelemetryPipelinesRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

