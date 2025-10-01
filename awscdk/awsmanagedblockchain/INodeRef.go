package awsmanagedblockchain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmanagedblockchain/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Node.
// Experimental.
type INodeRef interface {
	constructs.IConstruct
	// A reference to a Node resource.
	// Experimental.
	NodeRef() *NodeReference
}

// The jsii proxy for INodeRef
type jsiiProxy_INodeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INodeRef) NodeRef() *NodeReference {
	var returns *NodeReference
	_jsii_.Get(
		j,
		"nodeRef",
		&returns,
	)
	return returns
}

