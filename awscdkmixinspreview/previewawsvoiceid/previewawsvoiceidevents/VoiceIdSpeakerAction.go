package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.voiceid@VoiceIdSpeakerAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdSpeakerAction := awscdkmixinspreview.Events.NewVoiceIdSpeakerAction()
//
// Experimental.
type VoiceIdSpeakerAction interface {
}

// The jsii proxy struct for VoiceIdSpeakerAction
type jsiiProxy_VoiceIdSpeakerAction struct {
	_ byte // padding
}

// Experimental.
func NewVoiceIdSpeakerAction() VoiceIdSpeakerAction {
	_init_.Initialize()

	j := jsiiProxy_VoiceIdSpeakerAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSpeakerAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVoiceIdSpeakerAction_Override(v VoiceIdSpeakerAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSpeakerAction",
		nil, // no parameters
		v,
	)
}

// EventBridge event pattern for VoiceId Speaker Action.
// Experimental.
func VoiceIdSpeakerAction_EventPattern(options *VoiceIdSpeakerAction_VoiceIdSpeakerActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateVoiceIdSpeakerAction_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSpeakerAction",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

