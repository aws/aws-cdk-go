package interfacesawsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Endpoint.
// Experimental.
type IEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Endpoint resource.
	// Experimental.
	EndpointRef() *EndpointReference
}

// The jsii proxy for IEndpointRef
type jsiiProxy_IEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IEndpointRef) EndpointRef() *EndpointReference {
	var returns *EndpointReference
	_jsii_.Get(
		j,
		"endpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

