package previewawsmedialiveevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.medialive@MediaLiveChannelInputChange event types for Channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaLiveChannelInputChange := #error#.NewMediaLiveChannelInputChange()
//
// Experimental.
type ChannelEvents_MediaLiveChannelInputChange interface {
}

// The jsii proxy struct for ChannelEvents_MediaLiveChannelInputChange
type jsiiProxy_ChannelEvents_MediaLiveChannelInputChange struct {
	_ byte // padding
}

// Experimental.
func NewChannelEvents_MediaLiveChannelInputChange() ChannelEvents_MediaLiveChannelInputChange {
	_init_.Initialize()

	j := jsiiProxy_ChannelEvents_MediaLiveChannelInputChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.ChannelEvents.MediaLiveChannelInputChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewChannelEvents_MediaLiveChannelInputChange_Override(c ChannelEvents_MediaLiveChannelInputChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.ChannelEvents.MediaLiveChannelInputChange",
		nil, // no parameters
		c,
	)
}

