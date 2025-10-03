package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualNode.
// Experimental.
type IVirtualNodeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVirtualNodeRef
type jsiiProxy_IVirtualNodeRef struct {
	internal.Type__constructsIConstruct
}

