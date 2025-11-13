package interfacesawsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StackFleetAssociation.
// Experimental.
type IStackFleetAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StackFleetAssociation resource.
	// Experimental.
	StackFleetAssociationRef() *StackFleetAssociationReference
}

// The jsii proxy for IStackFleetAssociationRef
type jsiiProxy_IStackFleetAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStackFleetAssociationRef) StackFleetAssociationRef() *StackFleetAssociationReference {
	var returns *StackFleetAssociationReference
	_jsii_.Get(
		j,
		"stackFleetAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStackFleetAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStackFleetAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

