package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRegistration.
// Experimental.
type ITransitGatewayRegistrationRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayRegistration resource.
	// Experimental.
	TransitGatewayRegistrationRef() *TransitGatewayRegistrationReference
}

// The jsii proxy for ITransitGatewayRegistrationRef
type jsiiProxy_ITransitGatewayRegistrationRef struct {
	internal.Type__constructsIConstruct
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

