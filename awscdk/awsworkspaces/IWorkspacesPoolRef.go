package awsworkspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspaces/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkspacesPool.
// Experimental.
type IWorkspacesPoolRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWorkspacesPoolRef
type jsiiProxy_IWorkspacesPoolRef struct {
	internal.Type__constructsIConstruct
}

