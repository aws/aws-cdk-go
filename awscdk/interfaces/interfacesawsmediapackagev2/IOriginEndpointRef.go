package interfacesawsmediapackagev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OriginEndpoint.
// Experimental.
type IOriginEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OriginEndpoint resource.
	// Experimental.
	OriginEndpointRef() *OriginEndpointReference
}

// The jsii proxy for IOriginEndpointRef
type jsiiProxy_IOriginEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOriginEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOriginEndpointRef) OriginEndpointRef() *OriginEndpointReference {
	var returns *OriginEndpointReference
	_jsii_.Get(
		j,
		"originEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

