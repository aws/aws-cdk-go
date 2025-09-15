package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRouteTableAttachment.
// Experimental.
type ITransitGatewayRouteTableAttachmentRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayRouteTableAttachment resource.
	// Experimental.
	TransitGatewayRouteTableAttachmentRef() *TransitGatewayRouteTableAttachmentReference
}

// The jsii proxy for ITransitGatewayRouteTableAttachmentRef
type jsiiProxy_ITransitGatewayRouteTableAttachmentRef struct {
	internal.Type__constructsIConstruct
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

