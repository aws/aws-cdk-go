package awsorganizations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsorganizations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationalUnit.
// Experimental.
type IOrganizationalUnitRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOrganizationalUnitRef
type jsiiProxy_IOrganizationalUnitRef struct {
	internal.Type__constructsIConstruct
}

