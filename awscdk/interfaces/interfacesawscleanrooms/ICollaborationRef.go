package interfacesawscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Collaboration.
// Experimental.
type ICollaborationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Collaboration resource.
	// Experimental.
	CollaborationRef() *CollaborationReference
}

// The jsii proxy for ICollaborationRef
type jsiiProxy_ICollaborationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICollaborationRef) CollaborationRef() *CollaborationReference {
	var returns *CollaborationReference
	_jsii_.Get(
		j,
		"collaborationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICollaborationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICollaborationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

