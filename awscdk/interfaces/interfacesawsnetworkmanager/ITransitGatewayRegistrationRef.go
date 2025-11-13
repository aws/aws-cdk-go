package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRegistration.
// Experimental.
type ITransitGatewayRegistrationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayRegistration resource.
	// Experimental.
	TransitGatewayRegistrationRef() *TransitGatewayRegistrationReference
}

// The jsii proxy for ITransitGatewayRegistrationRef
type jsiiProxy_ITransitGatewayRegistrationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITransitGatewayRegistrationRef) TransitGatewayRegistrationRef() *TransitGatewayRegistrationReference {
	var returns *TransitGatewayRegistrationReference
	_jsii_.Get(
		j,
		"transitGatewayRegistrationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRegistrationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRegistrationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

