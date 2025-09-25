package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCGatewayAttachment.
// Experimental.
type IVPCGatewayAttachmentRef interface {
	constructs.IConstruct
	// A reference to a VPCGatewayAttachment resource.
	// Experimental.
	VpcGatewayAttachmentRef() *VPCGatewayAttachmentReference
}

// The jsii proxy for IVPCGatewayAttachmentRef
type jsiiProxy_IVPCGatewayAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCGatewayAttachmentRef) VpcGatewayAttachmentRef() *VPCGatewayAttachmentReference {
	var returns *VPCGatewayAttachmentReference
	_jsii_.Get(
		j,
		"vpcGatewayAttachmentRef",
		&returns,
	)
	return returns
}

