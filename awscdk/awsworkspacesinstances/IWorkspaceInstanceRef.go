package awsworkspacesinstances

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesinstances/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkspaceInstance.
// Experimental.
type IWorkspaceInstanceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWorkspaceInstanceRef
type jsiiProxy_IWorkspaceInstanceRef struct {
	internal.Type__constructsIConstruct
}

