package interfacesawsobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TelemetryEnrichment.
// Experimental.
type ITelemetryEnrichmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TelemetryEnrichment resource.
	// Experimental.
	TelemetryEnrichmentRef() *TelemetryEnrichmentReference
}

// The jsii proxy for ITelemetryEnrichmentRef
type jsiiProxy_ITelemetryEnrichmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITelemetryEnrichmentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITelemetryEnrichmentRef) TelemetryEnrichmentRef() *TelemetryEnrichmentReference {
	var returns *TelemetryEnrichmentReference
	_jsii_.Get(
		j,
		"telemetryEnrichmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITelemetryEnrichmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITelemetryEnrichmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

