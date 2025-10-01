package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectAttachment.
// Experimental.
type IConnectAttachmentRef interface {
	constructs.IConstruct
	// A reference to a ConnectAttachment resource.
	// Experimental.
	ConnectAttachmentRef() *ConnectAttachmentReference
}

// The jsii proxy for IConnectAttachmentRef
type jsiiProxy_IConnectAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConnectAttachmentRef) ConnectAttachmentRef() *ConnectAttachmentReference {
	var returns *ConnectAttachmentReference
	_jsii_.Get(
		j,
		"connectAttachmentRef",
		&returns,
	)
	return returns
}

