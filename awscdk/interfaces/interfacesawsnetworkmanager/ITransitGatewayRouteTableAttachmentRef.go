package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRouteTableAttachment.
// Experimental.
type ITransitGatewayRouteTableAttachmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayRouteTableAttachment resource.
	// Experimental.
	TransitGatewayRouteTableAttachmentRef() *TransitGatewayRouteTableAttachmentReference
}

// The jsii proxy for ITransitGatewayRouteTableAttachmentRef
type jsiiProxy_ITransitGatewayRouteTableAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITransitGatewayRouteTableAttachmentRef) TransitGatewayRouteTableAttachmentRef() *TransitGatewayRouteTableAttachmentReference {
	var returns *TransitGatewayRouteTableAttachmentReference
	_jsii_.Get(
		j,
		"transitGatewayRouteTableAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTableAttachmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTableAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

