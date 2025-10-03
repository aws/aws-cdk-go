package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcAttachment.
// Experimental.
type IVpcAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVpcAttachmentRef
type jsiiProxy_IVpcAttachmentRef struct {
	internal.Type__constructsIConstruct
}

