package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationalUnitAssociation.
// Experimental.
type IOrganizationalUnitAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a OrganizationalUnitAssociation resource.
	// Experimental.
	OrganizationalUnitAssociationRef() *OrganizationalUnitAssociationReference
}

// The jsii proxy for IOrganizationalUnitAssociationRef
type jsiiProxy_IOrganizationalUnitAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IOrganizationalUnitAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizationalUnitAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

