package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRegistration.
// Experimental.
type ITransitGatewayRegistrationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITransitGatewayRegistrationRef
type jsiiProxy_ITransitGatewayRegistrationRef struct {
	internal.Type__constructsIConstruct
}

