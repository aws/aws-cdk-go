package awsobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationCentralizationRule.
// Experimental.
type IOrganizationCentralizationRuleRef interface {
	constructs.IConstruct
	// A reference to a OrganizationCentralizationRule resource.
	// Experimental.
	OrganizationCentralizationRuleRef() *OrganizationCentralizationRuleReference
}

// The jsii proxy for IOrganizationCentralizationRuleRef
type jsiiProxy_IOrganizationCentralizationRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOrganizationCentralizationRuleRef) OrganizationCentralizationRuleRef() *OrganizationCentralizationRuleReference {
	var returns *OrganizationCentralizationRuleReference
	_jsii_.Get(
		j,
		"organizationCentralizationRuleRef",
		&returns,
	)
	return returns
}

