package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EndpointConfig.
// Experimental.
type IEndpointConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EndpointConfig resource.
	// Experimental.
	EndpointConfigRef() *EndpointConfigReference
}

// The jsii proxy for IEndpointConfigRef
type jsiiProxy_IEndpointConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IEndpointConfigRef) EndpointConfigRef() *EndpointConfigReference {
	var returns *EndpointConfigReference
	_jsii_.Get(
		j,
		"endpointConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

