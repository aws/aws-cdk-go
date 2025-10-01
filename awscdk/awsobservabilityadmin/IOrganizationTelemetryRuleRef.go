package awsobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationTelemetryRule.
// Experimental.
type IOrganizationTelemetryRuleRef interface {
	constructs.IConstruct
	// A reference to a OrganizationTelemetryRule resource.
	// Experimental.
	OrganizationTelemetryRuleRef() *OrganizationTelemetryRuleReference
}

// The jsii proxy for IOrganizationTelemetryRuleRef
type jsiiProxy_IOrganizationTelemetryRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOrganizationTelemetryRuleRef) OrganizationTelemetryRuleRef() *OrganizationTelemetryRuleReference {
	var returns *OrganizationTelemetryRuleReference
	_jsii_.Get(
		j,
		"organizationTelemetryRuleRef",
		&returns,
	)
	return returns
}

