package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Nodegroup.
// Experimental.
type INodegroupRef interface {
	constructs.IConstruct
	// A reference to a Nodegroup resource.
	// Experimental.
	NodegroupRef() *NodegroupReference
}

// The jsii proxy for INodegroupRef
type jsiiProxy_INodegroupRef struct {
	internal.Type__constructsIConstruct
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

