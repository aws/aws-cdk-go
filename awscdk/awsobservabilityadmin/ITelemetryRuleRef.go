package awsobservabilityadmin

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TelemetryRule.
// Experimental.
type ITelemetryRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITelemetryRuleRef
type jsiiProxy_ITelemetryRuleRef struct {
	internal.Type__constructsIConstruct
}

