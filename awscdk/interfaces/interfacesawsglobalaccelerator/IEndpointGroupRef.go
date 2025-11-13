package interfacesawsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EndpointGroup.
// Experimental.
type IEndpointGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EndpointGroup resource.
	// Experimental.
	EndpointGroupRef() *EndpointGroupReference
}

// The jsii proxy for IEndpointGroupRef
type jsiiProxy_IEndpointGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IEndpointGroupRef) EndpointGroupRef() *EndpointGroupReference {
	var returns *EndpointGroupReference
	_jsii_.Get(
		j,
		"endpointGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

