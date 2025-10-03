package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VoiceChannel.
// Experimental.
type IVoiceChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVoiceChannelRef
type jsiiProxy_IVoiceChannelRef struct {
	internal.Type__constructsIConstruct
}

