package interfacesawsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DestinationEndpoint.
// Experimental.
type IDestinationEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DestinationEndpoint resource.
	// Experimental.
	DestinationEndpointRef() *DestinationEndpointReference
}

// The jsii proxy for IDestinationEndpointRef
type jsiiProxy_IDestinationEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDestinationEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDestinationEndpointRef) DestinationEndpointRef() *DestinationEndpointReference {
	var returns *DestinationEndpointReference
	_jsii_.Get(
		j,
		"destinationEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDestinationEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDestinationEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

