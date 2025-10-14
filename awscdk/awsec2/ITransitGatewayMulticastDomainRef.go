package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayMulticastDomain.
// Experimental.
type ITransitGatewayMulticastDomainRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayMulticastDomain resource.
	// Experimental.
	TransitGatewayMulticastDomainRef() *TransitGatewayMulticastDomainReference
}

// The jsii proxy for ITransitGatewayMulticastDomainRef
type jsiiProxy_ITransitGatewayMulticastDomainRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransitGatewayMulticastDomainRef) TransitGatewayMulticastDomainRef() *TransitGatewayMulticastDomainReference {
	var returns *TransitGatewayMulticastDomainReference
	_jsii_.Get(
		j,
		"transitGatewayMulticastDomainRef",
		&returns,
	)
	return returns
}

