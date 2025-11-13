package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInterface.
// Experimental.
type INetworkInterfaceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NetworkInterface resource.
	// Experimental.
	NetworkInterfaceRef() *NetworkInterfaceReference
}

// The jsii proxy for INetworkInterfaceRef
type jsiiProxy_INetworkInterfaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_INetworkInterfaceRef) NetworkInterfaceRef() *NetworkInterfaceReference {
	var returns *NetworkInterfaceReference
	_jsii_.Get(
		j,
		"networkInterfaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkInterfaceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkInterfaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

