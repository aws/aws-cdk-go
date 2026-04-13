package previewawsmedialiveevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.medialive@MediaLiveMultiplexStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaLiveMultiplexStateChange := awscdkmixinspreview.Events.NewMediaLiveMultiplexStateChange()
//
// Experimental.
type MediaLiveMultiplexStateChange interface {
}

// The jsii proxy struct for MediaLiveMultiplexStateChange
type jsiiProxy_MediaLiveMultiplexStateChange struct {
	_ byte // padding
}

// Experimental.
func NewMediaLiveMultiplexStateChange() MediaLiveMultiplexStateChange {
	_init_.Initialize()

	j := jsiiProxy_MediaLiveMultiplexStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveMultiplexStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMediaLiveMultiplexStateChange_Override(m MediaLiveMultiplexStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveMultiplexStateChange",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MediaLive Multiplex State Change.
// Experimental.
func MediaLiveMultiplexStateChange_EventPattern(options *MediaLiveMultiplexStateChange_MediaLiveMultiplexStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMediaLiveMultiplexStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveMultiplexStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

