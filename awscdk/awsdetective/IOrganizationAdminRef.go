package awsdetective

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdetective/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationAdmin.
// Experimental.
type IOrganizationAdminRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOrganizationAdminRef
type jsiiProxy_IOrganizationAdminRef struct {
	internal.Type__constructsIConstruct
}

