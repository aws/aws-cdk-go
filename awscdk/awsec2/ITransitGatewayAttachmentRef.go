package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayAttachment.
// Experimental.
type ITransitGatewayAttachmentRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayAttachment resource.
	// Experimental.
	TransitGatewayAttachmentRef() *TransitGatewayAttachmentReference
}

// The jsii proxy for ITransitGatewayAttachmentRef
type jsiiProxy_ITransitGatewayAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransitGatewayAttachmentRef) TransitGatewayAttachmentRef() *TransitGatewayAttachmentReference {
	var returns *TransitGatewayAttachmentReference
	_jsii_.Get(
		j,
		"transitGatewayAttachmentRef",
		&returns,
	)
	return returns
}

