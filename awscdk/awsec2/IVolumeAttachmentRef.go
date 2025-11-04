package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VolumeAttachment.
// Experimental.
type IVolumeAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VolumeAttachment resource.
	// Experimental.
	VolumeAttachmentRef() *VolumeAttachmentReference
}

// The jsii proxy for IVolumeAttachmentRef
type jsiiProxy_IVolumeAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IVolumeAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVolumeAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

