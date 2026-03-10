package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.voiceid@VoiceIdBatchSpeakerEnrollmentAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdBatchSpeakerEnrollmentAction := awscdkmixinspreview.Events.NewVoiceIdBatchSpeakerEnrollmentAction()
//
// Experimental.
type VoiceIdBatchSpeakerEnrollmentAction interface {
}

// The jsii proxy struct for VoiceIdBatchSpeakerEnrollmentAction
type jsiiProxy_VoiceIdBatchSpeakerEnrollmentAction struct {
	_ byte // padding
}

// Experimental.
func NewVoiceIdBatchSpeakerEnrollmentAction() VoiceIdBatchSpeakerEnrollmentAction {
	_init_.Initialize()

	j := jsiiProxy_VoiceIdBatchSpeakerEnrollmentAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVoiceIdBatchSpeakerEnrollmentAction_Override(v VoiceIdBatchSpeakerEnrollmentAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction",
		nil, // no parameters
		v,
	)
}

// EventBridge event pattern for VoiceId Batch Speaker Enrollment Action.
// Experimental.
func VoiceIdBatchSpeakerEnrollmentAction_VoiceIdBatchSpeakerEnrollmentActionPattern(options *VoiceIdBatchSpeakerEnrollmentAction_VoiceIdBatchSpeakerEnrollmentActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateVoiceIdBatchSpeakerEnrollmentAction_VoiceIdBatchSpeakerEnrollmentActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchSpeakerEnrollmentAction",
		"voiceIdBatchSpeakerEnrollmentActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

