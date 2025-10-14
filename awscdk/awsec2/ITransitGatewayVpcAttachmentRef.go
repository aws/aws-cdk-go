package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayVpcAttachment.
// Experimental.
type ITransitGatewayVpcAttachmentRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayVpcAttachment resource.
	// Experimental.
	TransitGatewayVpcAttachmentRef() *TransitGatewayVpcAttachmentReference
}

// The jsii proxy for ITransitGatewayVpcAttachmentRef
type jsiiProxy_ITransitGatewayVpcAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransitGatewayVpcAttachmentRef) TransitGatewayVpcAttachmentRef() *TransitGatewayVpcAttachmentReference {
	var returns *TransitGatewayVpcAttachmentReference
	_jsii_.Get(
		j,
		"transitGatewayVpcAttachmentRef",
		&returns,
	)
	return returns
}

