package interfacesawsiotsitewise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Gateway.
// Experimental.
type IGatewayRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Gateway resource.
	// Experimental.
	GatewayRef() *GatewayReference
}

// The jsii proxy for IGatewayRef
type jsiiProxy_IGatewayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IGatewayRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGatewayRef) GatewayRef() *GatewayReference {
	var returns *GatewayReference
	_jsii_.Get(
		j,
		"gatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

