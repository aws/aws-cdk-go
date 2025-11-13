package interfacesawseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PodIdentityAssociation.
// Experimental.
type IPodIdentityAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PodIdentityAssociation resource.
	// Experimental.
	PodIdentityAssociationRef() *PodIdentityAssociationReference
}

// The jsii proxy for IPodIdentityAssociationRef
type jsiiProxy_IPodIdentityAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPodIdentityAssociationRef) PodIdentityAssociationRef() *PodIdentityAssociationReference {
	var returns *PodIdentityAssociationReference
	_jsii_.Get(
		j,
		"podIdentityAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPodIdentityAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPodIdentityAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

