package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayMulticastDomain.
// Experimental.
type ITransitGatewayMulticastDomainRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITransitGatewayMulticastDomainRef
type jsiiProxy_ITransitGatewayMulticastDomainRef struct {
	internal.Type__constructsIConstruct
}

