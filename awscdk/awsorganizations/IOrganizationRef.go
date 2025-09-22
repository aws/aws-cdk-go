package awsorganizations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsorganizations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Organization.
// Experimental.
type IOrganizationRef interface {
	constructs.IConstruct
	// A reference to a Organization resource.
	// Experimental.
	OrganizationRef() *OrganizationReference
}

// The jsii proxy for IOrganizationRef
type jsiiProxy_IOrganizationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOrganizationRef) OrganizationRef() *OrganizationReference {
	var returns *OrganizationReference
	_jsii_.Get(
		j,
		"organizationRef",
		&returns,
	)
	return returns
}

