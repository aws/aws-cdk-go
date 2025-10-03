package awsiottwinmaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workspace.
// Experimental.
type IWorkspaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWorkspaceRef
type jsiiProxy_IWorkspaceRef struct {
	internal.Type__constructsIConstruct
}

