package interfacesawsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Network.
// Experimental.
type INetworkRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Network resource.
	// Experimental.
	NetworkRef() *NetworkReference
}

// The jsii proxy for INetworkRef
type jsiiProxy_INetworkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_INetworkRef) NetworkRef() *NetworkReference {
	var returns *NetworkReference
	_jsii_.Get(
		j,
		"networkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

