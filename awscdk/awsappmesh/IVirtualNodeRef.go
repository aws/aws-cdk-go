package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualNode.
// Experimental.
type IVirtualNodeRef interface {
	constructs.IConstruct
	// A reference to a VirtualNode resource.
	// Experimental.
	VirtualNodeRef() *VirtualNodeReference
}

// The jsii proxy for IVirtualNodeRef
type jsiiProxy_IVirtualNodeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVirtualNodeRef) VirtualNodeRef() *VirtualNodeReference {
	var returns *VirtualNodeReference
	_jsii_.Get(
		j,
		"virtualNodeRef",
		&returns,
	)
	return returns
}

