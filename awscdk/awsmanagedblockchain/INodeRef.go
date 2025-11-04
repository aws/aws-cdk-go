package awsmanagedblockchain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmanagedblockchain/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Node.
// Experimental.
type INodeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Node resource.
	// Experimental.
	NodeRef() *NodeReference
}

// The jsii proxy for INodeRef
type jsiiProxy_INodeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_INodeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

