package interfacesawsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RuntimeEndpoint.
// Experimental.
type IRuntimeEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RuntimeEndpoint resource.
	// Experimental.
	RuntimeEndpointRef() *RuntimeEndpointReference
}

// The jsii proxy for IRuntimeEndpointRef
type jsiiProxy_IRuntimeEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRuntimeEndpointRef) RuntimeEndpointRef() *RuntimeEndpointReference {
	var returns *RuntimeEndpointReference
	_jsii_.Get(
		j,
		"runtimeEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

