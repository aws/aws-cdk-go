package previewawschimeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.chime@ChimeVoiceConnectorStreamingStatus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   chimeVoiceConnectorStreamingStatus := awscdkmixinspreview.Events.NewChimeVoiceConnectorStreamingStatus()
//
// Experimental.
type ChimeVoiceConnectorStreamingStatus interface {
}

// The jsii proxy struct for ChimeVoiceConnectorStreamingStatus
type jsiiProxy_ChimeVoiceConnectorStreamingStatus struct {
	_ byte // padding
}

// Experimental.
func NewChimeVoiceConnectorStreamingStatus() ChimeVoiceConnectorStreamingStatus {
	_init_.Initialize()

	j := jsiiProxy_ChimeVoiceConnectorStreamingStatus{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_chime.events.ChimeVoiceConnectorStreamingStatus",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewChimeVoiceConnectorStreamingStatus_Override(c ChimeVoiceConnectorStreamingStatus) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_chime.events.ChimeVoiceConnectorStreamingStatus",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Chime VoiceConnector Streaming Status.
// Experimental.
func ChimeVoiceConnectorStreamingStatus_EventPattern(options *ChimeVoiceConnectorStreamingStatus_ChimeVoiceConnectorStreamingStatusProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateChimeVoiceConnectorStreamingStatus_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_chime.events.ChimeVoiceConnectorStreamingStatus",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

