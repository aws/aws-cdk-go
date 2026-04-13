package previewawsmedialiveevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.medialive@MediaLiveChannelInputChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaLiveChannelInputChange := awscdkmixinspreview.Events.NewMediaLiveChannelInputChange()
//
// Experimental.
type MediaLiveChannelInputChange interface {
}

// The jsii proxy struct for MediaLiveChannelInputChange
type jsiiProxy_MediaLiveChannelInputChange struct {
	_ byte // padding
}

// Experimental.
func NewMediaLiveChannelInputChange() MediaLiveChannelInputChange {
	_init_.Initialize()

	j := jsiiProxy_MediaLiveChannelInputChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelInputChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMediaLiveChannelInputChange_Override(m MediaLiveChannelInputChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelInputChange",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MediaLive Channel Input Change.
// Experimental.
func MediaLiveChannelInputChange_EventPattern(options *MediaLiveChannelInputChange_MediaLiveChannelInputChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMediaLiveChannelInputChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelInputChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

