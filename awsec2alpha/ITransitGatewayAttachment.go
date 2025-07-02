package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
)

// Represents a Transit Gateway Attachment.
// Experimental.
type ITransitGatewayAttachment interface {
	awscdk.IResource
	// The ID of the transit gateway attachment.
	// Experimental.
	TransitGatewayAttachmentId() *string
}

// The jsii proxy for ITransitGatewayAttachment
type jsiiProxy_ITransitGatewayAttachment struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITransitGatewayAttachment) TransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentId",
		&returns,
	)
	return returns
}

