package interfacesawsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StackUserAssociation.
// Experimental.
type IStackUserAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StackUserAssociation resource.
	// Experimental.
	StackUserAssociationRef() *StackUserAssociationReference
}

// The jsii proxy for IStackUserAssociationRef
type jsiiProxy_IStackUserAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStackUserAssociationRef) StackUserAssociationRef() *StackUserAssociationReference {
	var returns *StackUserAssociationReference
	_jsii_.Get(
		j,
		"stackUserAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStackUserAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStackUserAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

