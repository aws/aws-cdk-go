package interfacesawsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationFleetAssociation.
// Experimental.
type IApplicationFleetAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ApplicationFleetAssociation resource.
	// Experimental.
	ApplicationFleetAssociationRef() *ApplicationFleetAssociationReference
}

// The jsii proxy for IApplicationFleetAssociationRef
type jsiiProxy_IApplicationFleetAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IApplicationFleetAssociationRef) ApplicationFleetAssociationRef() *ApplicationFleetAssociationReference {
	var returns *ApplicationFleetAssociationReference
	_jsii_.Get(
		j,
		"applicationFleetAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationFleetAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationFleetAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

