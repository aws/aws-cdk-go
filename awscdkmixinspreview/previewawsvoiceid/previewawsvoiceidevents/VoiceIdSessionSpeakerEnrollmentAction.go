package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.voiceid@VoiceIdSessionSpeakerEnrollmentAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdSessionSpeakerEnrollmentAction := awscdkmixinspreview.Events.NewVoiceIdSessionSpeakerEnrollmentAction()
//
// Experimental.
type VoiceIdSessionSpeakerEnrollmentAction interface {
}

// The jsii proxy struct for VoiceIdSessionSpeakerEnrollmentAction
type jsiiProxy_VoiceIdSessionSpeakerEnrollmentAction struct {
	_ byte // padding
}

// Experimental.
func NewVoiceIdSessionSpeakerEnrollmentAction() VoiceIdSessionSpeakerEnrollmentAction {
	_init_.Initialize()

	j := jsiiProxy_VoiceIdSessionSpeakerEnrollmentAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSessionSpeakerEnrollmentAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVoiceIdSessionSpeakerEnrollmentAction_Override(v VoiceIdSessionSpeakerEnrollmentAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSessionSpeakerEnrollmentAction",
		nil, // no parameters
		v,
	)
}

// EventBridge event pattern for VoiceId Session Speaker Enrollment Action.
// Experimental.
func VoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionPattern(options *VoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateVoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdSessionSpeakerEnrollmentAction",
		"voiceIdSessionSpeakerEnrollmentActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

