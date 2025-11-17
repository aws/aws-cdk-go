package interfacesawsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouterNetworkInterface.
// Experimental.
type IRouterNetworkInterfaceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RouterNetworkInterface resource.
	// Experimental.
	RouterNetworkInterfaceRef() *RouterNetworkInterfaceReference
}

// The jsii proxy for IRouterNetworkInterfaceRef
type jsiiProxy_IRouterNetworkInterfaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRouterNetworkInterfaceRef) RouterNetworkInterfaceRef() *RouterNetworkInterfaceReference {
	var returns *RouterNetworkInterfaceReference
	_jsii_.Get(
		j,
		"routerNetworkInterfaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterNetworkInterfaceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterNetworkInterfaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

