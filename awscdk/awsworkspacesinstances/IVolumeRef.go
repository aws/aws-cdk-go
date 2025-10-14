package awsworkspacesinstances

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesinstances/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Volume.
// Experimental.
type IVolumeRef interface {
	constructs.IConstruct
	// A reference to a Volume resource.
	// Experimental.
	VolumeRef() *VolumeReference
}

// The jsii proxy for IVolumeRef
type jsiiProxy_IVolumeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVolumeRef) VolumeRef() *VolumeReference {
	var returns *VolumeReference
	_jsii_.Get(
		j,
		"volumeRef",
		&returns,
	)
	return returns
}

