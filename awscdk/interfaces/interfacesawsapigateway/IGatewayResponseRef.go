package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GatewayResponse.
// Experimental.
type IGatewayResponseRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GatewayResponse resource.
	// Experimental.
	GatewayResponseRef() *GatewayResponseReference
}

// The jsii proxy for IGatewayResponseRef
type jsiiProxy_IGatewayResponseRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IGatewayResponseRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGatewayResponseRef) GatewayResponseRef() *GatewayResponseReference {
	var returns *GatewayResponseReference
	_jsii_.Get(
		j,
		"gatewayResponseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayResponseRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayResponseRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

