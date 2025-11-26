package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.voiceid@VoiceIdSessionSpeakerEnrollmentAction event types for Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdSessionSpeakerEnrollmentAction := #error#.NewVoiceIdSessionSpeakerEnrollmentAction()
//
// Experimental.
type DomainEvents_VoiceIdSessionSpeakerEnrollmentAction interface {
}

// The jsii proxy struct for DomainEvents_VoiceIdSessionSpeakerEnrollmentAction
type jsiiProxy_DomainEvents_VoiceIdSessionSpeakerEnrollmentAction struct {
	_ byte // padding
}

// Experimental.
func NewDomainEvents_VoiceIdSessionSpeakerEnrollmentAction() DomainEvents_VoiceIdSessionSpeakerEnrollmentAction {
	_init_.Initialize()

	j := jsiiProxy_DomainEvents_VoiceIdSessionSpeakerEnrollmentAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSessionSpeakerEnrollmentAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDomainEvents_VoiceIdSessionSpeakerEnrollmentAction_Override(d DomainEvents_VoiceIdSessionSpeakerEnrollmentAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdSessionSpeakerEnrollmentAction",
		nil, // no parameters
		d,
	)
}

