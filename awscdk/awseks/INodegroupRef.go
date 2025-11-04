package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Nodegroup.
// Experimental.
type INodegroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Nodegroup resource.
	// Experimental.
	NodegroupRef() *NodegroupReference
}

// The jsii proxy for INodegroupRef
type jsiiProxy_INodegroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_INodegroupRef) NodegroupRef() *NodegroupReference {
	var returns *NodegroupReference
	_jsii_.Get(
		j,
		"nodegroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodegroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodegroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

