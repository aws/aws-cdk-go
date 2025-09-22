package awsmediatailor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlaybackConfiguration.
// Experimental.
type IPlaybackConfigurationRef interface {
	constructs.IConstruct
	// A reference to a PlaybackConfiguration resource.
	// Experimental.
	PlaybackConfigurationRef() *PlaybackConfigurationReference
}

// The jsii proxy for IPlaybackConfigurationRef
type jsiiProxy_IPlaybackConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPlaybackConfigurationRef) PlaybackConfigurationRef() *PlaybackConfigurationReference {
	var returns *PlaybackConfigurationReference
	_jsii_.Get(
		j,
		"playbackConfigurationRef",
		&returns,
	)
	return returns
}

