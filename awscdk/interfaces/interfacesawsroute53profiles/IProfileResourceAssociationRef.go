package interfacesawsroute53profiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53profiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfileResourceAssociation.
// Experimental.
type IProfileResourceAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ProfileResourceAssociation resource.
	// Experimental.
	ProfileResourceAssociationRef() *ProfileResourceAssociationReference
}

// The jsii proxy for IProfileResourceAssociationRef
type jsiiProxy_IProfileResourceAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IProfileResourceAssociationRef) ProfileResourceAssociationRef() *ProfileResourceAssociationReference {
	var returns *ProfileResourceAssociationReference
	_jsii_.Get(
		j,
		"profileResourceAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfileResourceAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfileResourceAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

