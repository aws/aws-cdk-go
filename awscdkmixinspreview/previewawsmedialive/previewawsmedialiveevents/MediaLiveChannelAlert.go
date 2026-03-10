package previewawsmedialiveevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.medialive@MediaLiveChannelAlert.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaLiveChannelAlert := awscdkmixinspreview.Events.NewMediaLiveChannelAlert()
//
// Experimental.
type MediaLiveChannelAlert interface {
}

// The jsii proxy struct for MediaLiveChannelAlert
type jsiiProxy_MediaLiveChannelAlert struct {
	_ byte // padding
}

// Experimental.
func NewMediaLiveChannelAlert() MediaLiveChannelAlert {
	_init_.Initialize()

	j := jsiiProxy_MediaLiveChannelAlert{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelAlert",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMediaLiveChannelAlert_Override(m MediaLiveChannelAlert) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelAlert",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MediaLive Channel Alert.
// Experimental.
func MediaLiveChannelAlert_MediaLiveChannelAlertPattern(options *MediaLiveChannelAlert_MediaLiveChannelAlertProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMediaLiveChannelAlert_MediaLiveChannelAlertPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_medialive.events.MediaLiveChannelAlert",
		"mediaLiveChannelAlertPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

