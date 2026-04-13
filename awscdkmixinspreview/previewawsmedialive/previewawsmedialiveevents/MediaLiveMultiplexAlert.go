package previewawsmedialiveevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.medialive@MediaLiveMultiplexAlert.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaLiveMultiplexAlert := awscdkmixinspreview.Events.NewMediaLiveMultiplexAlert()
//
// Experimental.
type MediaLiveMultiplexAlert interface {
}

// The jsii proxy struct for MediaLiveMultiplexAlert
type jsiiProxy_MediaLiveMultiplexAlert struct {
	_ byte // padding
}

// Experimental.
func NewMediaLiveMultiplexAlert() MediaLiveMultiplexAlert {
	_init_.Initialize()

	j := jsiiProxy_MediaLiveMultiplexAlert{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveMultiplexAlert",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMediaLiveMultiplexAlert_Override(m MediaLiveMultiplexAlert) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveMultiplexAlert",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MediaLive Multiplex Alert.
// Experimental.
func MediaLiveMultiplexAlert_EventPattern(options *MediaLiveMultiplexAlert_MediaLiveMultiplexAlertProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMediaLiveMultiplexAlert_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveMultiplexAlert",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

