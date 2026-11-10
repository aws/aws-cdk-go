package interfacesawsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GatewayClusterAssociations.
// Experimental.
type IGatewayClusterAssociationsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GatewayClusterAssociations resource.
	// Experimental.
	GatewayClusterAssociationsRef() *GatewayClusterAssociationsReference
}

// The jsii proxy for IGatewayClusterAssociationsRef
type jsiiProxy_IGatewayClusterAssociationsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IGatewayClusterAssociationsRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGatewayClusterAssociationsRef) GatewayClusterAssociationsRef() *GatewayClusterAssociationsReference {
	var returns *GatewayClusterAssociationsReference
	_jsii_.Get(
		j,
		"gatewayClusterAssociationsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayClusterAssociationsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayClusterAssociationsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

