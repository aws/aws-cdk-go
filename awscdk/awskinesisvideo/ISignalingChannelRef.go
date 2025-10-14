package awskinesisvideo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisvideo/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SignalingChannel.
// Experimental.
type ISignalingChannelRef interface {
	constructs.IConstruct
	// A reference to a SignalingChannel resource.
	// Experimental.
	SignalingChannelRef() *SignalingChannelReference
}

// The jsii proxy for ISignalingChannelRef
type jsiiProxy_ISignalingChannelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISignalingChannelRef) SignalingChannelRef() *SignalingChannelReference {
	var returns *SignalingChannelReference
	_jsii_.Get(
		j,
		"signalingChannelRef",
		&returns,
	)
	return returns
}

