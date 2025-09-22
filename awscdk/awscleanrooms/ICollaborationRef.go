package awscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Collaboration.
// Experimental.
type ICollaborationRef interface {
	constructs.IConstruct
	// A reference to a Collaboration resource.
	// Experimental.
	CollaborationRef() *CollaborationReference
}

// The jsii proxy for ICollaborationRef
type jsiiProxy_ICollaborationRef struct {
	internal.Type__constructsIConstruct
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

