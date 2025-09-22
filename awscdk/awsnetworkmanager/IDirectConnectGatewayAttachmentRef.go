package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectConnectGatewayAttachment.
// Experimental.
type IDirectConnectGatewayAttachmentRef interface {
	constructs.IConstruct
	// A reference to a DirectConnectGatewayAttachment resource.
	// Experimental.
	DirectConnectGatewayAttachmentRef() *DirectConnectGatewayAttachmentReference
}

// The jsii proxy for IDirectConnectGatewayAttachmentRef
type jsiiProxy_IDirectConnectGatewayAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDirectConnectGatewayAttachmentRef) DirectConnectGatewayAttachmentRef() *DirectConnectGatewayAttachmentReference {
	var returns *DirectConnectGatewayAttachmentReference
	_jsii_.Get(
		j,
		"directConnectGatewayAttachmentRef",
		&returns,
	)
	return returns
}

