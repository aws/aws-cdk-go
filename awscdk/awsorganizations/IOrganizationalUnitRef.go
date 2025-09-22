package awsorganizations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsorganizations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationalUnit.
// Experimental.
type IOrganizationalUnitRef interface {
	constructs.IConstruct
	// A reference to a OrganizationalUnit resource.
	// Experimental.
	OrganizationalUnitRef() *OrganizationalUnitReference
}

// The jsii proxy for IOrganizationalUnitRef
type jsiiProxy_IOrganizationalUnitRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOrganizationalUnitRef) OrganizationalUnitRef() *OrganizationalUnitReference {
	var returns *OrganizationalUnitReference
	_jsii_.Get(
		j,
		"organizationalUnitRef",
		&returns,
	)
	return returns
}

