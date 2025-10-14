package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayMulticastDomainAssociation.
// Experimental.
type ITransitGatewayMulticastDomainAssociationRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayMulticastDomainAssociation resource.
	// Experimental.
	TransitGatewayMulticastDomainAssociationRef() *TransitGatewayMulticastDomainAssociationReference
}

// The jsii proxy for ITransitGatewayMulticastDomainAssociationRef
type jsiiProxy_ITransitGatewayMulticastDomainAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransitGatewayMulticastDomainAssociationRef) TransitGatewayMulticastDomainAssociationRef() *TransitGatewayMulticastDomainAssociationReference {
	var returns *TransitGatewayMulticastDomainAssociationReference
	_jsii_.Get(
		j,
		"transitGatewayMulticastDomainAssociationRef",
		&returns,
	)
	return returns
}

