package interfacesawslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkConnector.
// Experimental.
type INetworkConnectorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NetworkConnector resource.
	// Experimental.
	NetworkConnectorRef() *NetworkConnectorReference
}

// The jsii proxy for INetworkConnectorRef
type jsiiProxy_INetworkConnectorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_INetworkConnectorRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INetworkConnectorRef) NetworkConnectorRef() *NetworkConnectorReference {
	var returns *NetworkConnectorReference
	_jsii_.Get(
		j,
		"networkConnectorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkConnectorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkConnectorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

