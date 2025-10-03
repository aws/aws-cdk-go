package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationConfiguration.
// Experimental.
type IOrganizationConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOrganizationConfigurationRef
type jsiiProxy_IOrganizationConfigurationRef struct {
	internal.Type__constructsIConstruct
}

