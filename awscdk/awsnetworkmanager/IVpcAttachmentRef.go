package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcAttachment.
// Experimental.
type IVpcAttachmentRef interface {
	constructs.IConstruct
	// A reference to a VpcAttachment resource.
	// Experimental.
	VpcAttachmentRef() *VpcAttachmentReference
}

// The jsii proxy for IVpcAttachmentRef
type jsiiProxy_IVpcAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVpcAttachmentRef) VpcAttachmentRef() *VpcAttachmentReference {
	var returns *VpcAttachmentReference
	_jsii_.Get(
		j,
		"vpcAttachmentRef",
		&returns,
	)
	return returns
}

