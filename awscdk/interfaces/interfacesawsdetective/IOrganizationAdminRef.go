package interfacesawsdetective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdetective/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationAdmin.
// Experimental.
type IOrganizationAdminRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OrganizationAdmin resource.
	// Experimental.
	OrganizationAdminRef() *OrganizationAdminReference
}

// The jsii proxy for IOrganizationAdminRef
type jsiiProxy_IOrganizationAdminRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IOrganizationAdminRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

