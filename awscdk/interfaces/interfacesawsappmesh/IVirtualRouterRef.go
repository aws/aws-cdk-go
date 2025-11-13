package interfacesawsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualRouter.
// Experimental.
type IVirtualRouterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VirtualRouter resource.
	// Experimental.
	VirtualRouterRef() *VirtualRouterReference
}

// The jsii proxy for IVirtualRouterRef
type jsiiProxy_IVirtualRouterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVirtualRouterRef) VirtualRouterRef() *VirtualRouterReference {
	var returns *VirtualRouterReference
	_jsii_.Get(
		j,
		"virtualRouterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualRouterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualRouterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

