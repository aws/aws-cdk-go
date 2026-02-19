package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceConnectEndpoint.
// Experimental.
type IInstanceConnectEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InstanceConnectEndpoint resource.
	// Experimental.
	InstanceConnectEndpointRef() *InstanceConnectEndpointReference
}

// The jsii proxy for IInstanceConnectEndpointRef
type jsiiProxy_IInstanceConnectEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IInstanceConnectEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IInstanceConnectEndpointRef) InstanceConnectEndpointRef() *InstanceConnectEndpointReference {
	var returns *InstanceConnectEndpointReference
	_jsii_.Get(
		j,
		"instanceConnectEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceConnectEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceConnectEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

