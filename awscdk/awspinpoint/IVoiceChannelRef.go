package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VoiceChannel.
// Experimental.
type IVoiceChannelRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VoiceChannel resource.
	// Experimental.
	VoiceChannelRef() *VoiceChannelReference
}

// The jsii proxy for IVoiceChannelRef
type jsiiProxy_IVoiceChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IVoiceChannelRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVoiceChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

