package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayMulticastDomainAssociation.
// Experimental.
type ITransitGatewayMulticastDomainAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayMulticastDomainAssociation resource.
	// Experimental.
	TransitGatewayMulticastDomainAssociationRef() *TransitGatewayMulticastDomainAssociationReference
}

// The jsii proxy for ITransitGatewayMulticastDomainAssociationRef
type jsiiProxy_ITransitGatewayMulticastDomainAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ITransitGatewayMulticastDomainAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayMulticastDomainAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

