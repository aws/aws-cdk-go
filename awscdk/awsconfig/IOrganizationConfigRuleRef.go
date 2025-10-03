package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationConfigRule.
// Experimental.
type IOrganizationConfigRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOrganizationConfigRuleRef
type jsiiProxy_IOrganizationConfigRuleRef struct {
	internal.Type__constructsIConstruct
}

