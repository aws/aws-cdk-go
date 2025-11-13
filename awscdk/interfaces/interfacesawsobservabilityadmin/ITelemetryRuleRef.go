package interfacesawsobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TelemetryRule.
// Experimental.
type ITelemetryRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TelemetryRule resource.
	// Experimental.
	TelemetryRuleRef() *TelemetryRuleReference
}

// The jsii proxy for ITelemetryRuleRef
type jsiiProxy_ITelemetryRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITelemetryRuleRef) TelemetryRuleRef() *TelemetryRuleReference {
	var returns *TelemetryRuleReference
	_jsii_.Get(
		j,
		"telemetryRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITelemetryRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITelemetryRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

