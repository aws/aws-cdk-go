package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationConfigRule.
// Experimental.
type IOrganizationConfigRuleRef interface {
	constructs.IConstruct
	// A reference to a OrganizationConfigRule resource.
	// Experimental.
	OrganizationConfigRuleRef() *OrganizationConfigRuleReference
}

// The jsii proxy for IOrganizationConfigRuleRef
type jsiiProxy_IOrganizationConfigRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOrganizationConfigRuleRef) OrganizationConfigRuleRef() *OrganizationConfigRuleReference {
	var returns *OrganizationConfigRuleReference
	_jsii_.Get(
		j,
		"organizationConfigRuleRef",
		&returns,
	)
	return returns
}

