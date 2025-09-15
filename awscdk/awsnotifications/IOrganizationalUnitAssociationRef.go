package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationalUnitAssociation.
// Experimental.
type IOrganizationalUnitAssociationRef interface {
	constructs.IConstruct
	// A reference to a OrganizationalUnitAssociation resource.
	// Experimental.
	OrganizationalUnitAssociationRef() *OrganizationalUnitAssociationReference
}

// The jsii proxy for IOrganizationalUnitAssociationRef
type jsiiProxy_IOrganizationalUnitAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOrganizationalUnitAssociationRef) OrganizationalUnitAssociationRef() *OrganizationalUnitAssociationReference {
	var returns *OrganizationalUnitAssociationReference
	_jsii_.Get(
		j,
		"organizationalUnitAssociationRef",
		&returns,
	)
	return returns
}

