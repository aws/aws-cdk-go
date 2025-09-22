package awsdetective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdetective/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationAdmin.
// Experimental.
type IOrganizationAdminRef interface {
	constructs.IConstruct
	// A reference to a OrganizationAdmin resource.
	// Experimental.
	OrganizationAdminRef() *OrganizationAdminReference
}

// The jsii proxy for IOrganizationAdminRef
type jsiiProxy_IOrganizationAdminRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOrganizationAdminRef) OrganizationAdminRef() *OrganizationAdminReference {
	var returns *OrganizationAdminReference
	_jsii_.Get(
		j,
		"organizationAdminRef",
		&returns,
	)
	return returns
}

