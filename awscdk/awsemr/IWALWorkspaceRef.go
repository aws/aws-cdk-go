package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WALWorkspace.
// Experimental.
type IWALWorkspaceRef interface {
	constructs.IConstruct
	// A reference to a WALWorkspace resource.
	// Experimental.
	WalWorkspaceRef() *WALWorkspaceReference
}

// The jsii proxy for IWALWorkspaceRef
type jsiiProxy_IWALWorkspaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWALWorkspaceRef) WalWorkspaceRef() *WALWorkspaceReference {
	var returns *WALWorkspaceReference
	_jsii_.Get(
		j,
		"walWorkspaceRef",
		&returns,
	)
	return returns
}

