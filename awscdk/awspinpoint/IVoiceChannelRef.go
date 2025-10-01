package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VoiceChannel.
// Experimental.
type IVoiceChannelRef interface {
	constructs.IConstruct
	// A reference to a VoiceChannel resource.
	// Experimental.
	VoiceChannelRef() *VoiceChannelReference
}

// The jsii proxy for IVoiceChannelRef
type jsiiProxy_IVoiceChannelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVoiceChannelRef) VoiceChannelRef() *VoiceChannelReference {
	var returns *VoiceChannelReference
	_jsii_.Get(
		j,
		"voiceChannelRef",
		&returns,
	)
	return returns
}

