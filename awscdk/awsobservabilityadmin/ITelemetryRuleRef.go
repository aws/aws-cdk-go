package awsobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TelemetryRule.
// Experimental.
type ITelemetryRuleRef interface {
	constructs.IConstruct
	// A reference to a TelemetryRule resource.
	// Experimental.
	TelemetryRuleRef() *TelemetryRuleReference
}

// The jsii proxy for ITelemetryRuleRef
type jsiiProxy_ITelemetryRuleRef struct {
	internal.Type__constructsIConstruct
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

