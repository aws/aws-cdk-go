package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LinkAssociation.
// Experimental.
type ILinkAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LinkAssociation resource.
	// Experimental.
	LinkAssociationRef() *LinkAssociationReference
}

// The jsii proxy for ILinkAssociationRef
type jsiiProxy_ILinkAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILinkAssociationRef) LinkAssociationRef() *LinkAssociationReference {
	var returns *LinkAssociationReference
	_jsii_.Get(
		j,
		"linkAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILinkAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILinkAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

