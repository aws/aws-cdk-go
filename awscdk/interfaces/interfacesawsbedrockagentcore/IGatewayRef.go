package interfacesawsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore/internal"
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

