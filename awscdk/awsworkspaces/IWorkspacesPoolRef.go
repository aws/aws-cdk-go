package awsworkspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspaces/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkspacesPool.
// Experimental.
type IWorkspacesPoolRef interface {
	constructs.IConstruct
	// A reference to a WorkspacesPool resource.
	// Experimental.
	WorkspacesPoolRef() *WorkspacesPoolReference
}

// The jsii proxy for IWorkspacesPoolRef
type jsiiProxy_IWorkspacesPoolRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWorkspacesPoolRef) WorkspacesPoolRef() *WorkspacesPoolReference {
	var returns *WorkspacesPoolReference
	_jsii_.Get(
		j,
		"workspacesPoolRef",
		&returns,
	)
	return returns
}

