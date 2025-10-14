package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInterfaceAttachment.
// Experimental.
type INetworkInterfaceAttachmentRef interface {
	constructs.IConstruct
	// A reference to a NetworkInterfaceAttachment resource.
	// Experimental.
	NetworkInterfaceAttachmentRef() *NetworkInterfaceAttachmentReference
}

// The jsii proxy for INetworkInterfaceAttachmentRef
type jsiiProxy_INetworkInterfaceAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkInterfaceAttachmentRef) NetworkInterfaceAttachmentRef() *NetworkInterfaceAttachmentReference {
	var returns *NetworkInterfaceAttachmentReference
	_jsii_.Get(
		j,
		"networkInterfaceAttachmentRef",
		&returns,
	)
	return returns
}

