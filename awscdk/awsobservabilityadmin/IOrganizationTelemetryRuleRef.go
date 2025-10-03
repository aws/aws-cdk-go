package awsobservabilityadmin

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationTelemetryRule.
// Experimental.
type IOrganizationTelemetryRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOrganizationTelemetryRuleRef
type jsiiProxy_IOrganizationTelemetryRuleRef struct {
	internal.Type__constructsIConstruct
}

