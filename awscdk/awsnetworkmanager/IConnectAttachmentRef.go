package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectAttachment.
// Experimental.
type IConnectAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConnectAttachmentRef
type jsiiProxy_IConnectAttachmentRef struct {
	internal.Type__constructsIConstruct
}

