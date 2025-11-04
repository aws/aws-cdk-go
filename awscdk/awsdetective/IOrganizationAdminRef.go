package awsdetective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdetective/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationAdmin.
// Experimental.
type IOrganizationAdminRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a OrganizationAdmin resource.
	// Experimental.
	OrganizationAdminRef() *OrganizationAdminReference
}

// The jsii proxy for IOrganizationAdminRef
type jsiiProxy_IOrganizationAdminRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IOrganizationAdminRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizationAdminRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

