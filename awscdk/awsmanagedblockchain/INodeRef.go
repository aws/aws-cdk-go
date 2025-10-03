package awsmanagedblockchain

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmanagedblockchain/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Node.
// Experimental.
type INodeRef interface {
	constructs.IConstruct
}

// The jsii proxy for INodeRef
type jsiiProxy_INodeRef struct {
	internal.Type__constructsIConstruct
}

