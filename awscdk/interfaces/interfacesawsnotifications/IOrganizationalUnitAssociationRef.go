package interfacesawsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationalUnitAssociation.
// Experimental.
type IOrganizationalUnitAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OrganizationalUnitAssociation resource.
	// Experimental.
	OrganizationalUnitAssociationRef() *OrganizationalUnitAssociationReference
}

// The jsii proxy for IOrganizationalUnitAssociationRef
type jsiiProxy_IOrganizationalUnitAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IOrganizationalUnitAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

