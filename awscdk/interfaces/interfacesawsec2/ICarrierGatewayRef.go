package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CarrierGateway.
// Experimental.
type ICarrierGatewayRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CarrierGateway resource.
	// Experimental.
	CarrierGatewayRef() *CarrierGatewayReference
}

// The jsii proxy for ICarrierGatewayRef
type jsiiProxy_ICarrierGatewayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICarrierGatewayRef) CarrierGatewayRef() *CarrierGatewayReference {
	var returns *CarrierGatewayReference
	_jsii_.Get(
		j,
		"carrierGatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICarrierGatewayRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICarrierGatewayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

