package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRouteTableAttachment.
// Experimental.
type ITransitGatewayRouteTableAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITransitGatewayRouteTableAttachmentRef
type jsiiProxy_ITransitGatewayRouteTableAttachmentRef struct {
	internal.Type__constructsIConstruct
}

