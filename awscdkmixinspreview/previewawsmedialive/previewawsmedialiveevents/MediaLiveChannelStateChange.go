package previewawsmedialiveevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.medialive@MediaLiveChannelStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaLiveChannelStateChange := awscdkmixinspreview.Events.NewMediaLiveChannelStateChange()
//
// Experimental.
type MediaLiveChannelStateChange interface {
}

// The jsii proxy struct for MediaLiveChannelStateChange
type jsiiProxy_MediaLiveChannelStateChange struct {
	_ byte // padding
}

// Experimental.
func NewMediaLiveChannelStateChange() MediaLiveChannelStateChange {
	_init_.Initialize()

	j := jsiiProxy_MediaLiveChannelStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMediaLiveChannelStateChange_Override(m MediaLiveChannelStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelStateChange",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MediaLive Channel State Change.
// Experimental.
func MediaLiveChannelStateChange_EventPattern(options *MediaLiveChannelStateChange_MediaLiveChannelStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMediaLiveChannelStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

