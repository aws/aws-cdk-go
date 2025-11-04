package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualNode.
// Experimental.
type IVirtualNodeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VirtualNode resource.
	// Experimental.
	VirtualNodeRef() *VirtualNodeReference
}

// The jsii proxy for IVirtualNodeRef
type jsiiProxy_IVirtualNodeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IVirtualNodeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualNodeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

