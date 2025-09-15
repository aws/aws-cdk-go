package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VolumeAttachment.
// Experimental.
type IVolumeAttachmentRef interface {
	constructs.IConstruct
	// A reference to a VolumeAttachment resource.
	// Experimental.
	VolumeAttachmentRef() *VolumeAttachmentReference
}

// The jsii proxy for IVolumeAttachmentRef
type jsiiProxy_IVolumeAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVolumeAttachmentRef) VolumeAttachmentRef() *VolumeAttachmentReference {
	var returns *VolumeAttachmentReference
	_jsii_.Get(
		j,
		"volumeAttachmentRef",
		&returns,
	)
	return returns
}

