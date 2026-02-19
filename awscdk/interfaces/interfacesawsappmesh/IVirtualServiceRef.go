package interfacesawsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualService.
// Experimental.
type IVirtualServiceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VirtualService resource.
	// Experimental.
	VirtualServiceRef() *VirtualServiceReference
}

// The jsii proxy for IVirtualServiceRef
type jsiiProxy_IVirtualServiceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVirtualServiceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVirtualServiceRef) VirtualServiceRef() *VirtualServiceReference {
	var returns *VirtualServiceReference
	_jsii_.Get(
		j,
		"virtualServiceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualServiceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualServiceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

