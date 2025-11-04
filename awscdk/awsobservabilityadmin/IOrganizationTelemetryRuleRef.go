package awsobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationTelemetryRule.
// Experimental.
type IOrganizationTelemetryRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a OrganizationTelemetryRule resource.
	// Experimental.
	OrganizationTelemetryRuleRef() *OrganizationTelemetryRuleReference
}

// The jsii proxy for IOrganizationTelemetryRuleRef
type jsiiProxy_IOrganizationTelemetryRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IOrganizationTelemetryRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizationTelemetryRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

