package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.voiceid@VoiceIdEvaluateSessionAction event types for Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdEvaluateSessionAction := #error#.NewVoiceIdEvaluateSessionAction()
//
// Experimental.
type DomainEvents_VoiceIdEvaluateSessionAction interface {
}

// The jsii proxy struct for DomainEvents_VoiceIdEvaluateSessionAction
type jsiiProxy_DomainEvents_VoiceIdEvaluateSessionAction struct {
	_ byte // padding
}

// Experimental.
func NewDomainEvents_VoiceIdEvaluateSessionAction() DomainEvents_VoiceIdEvaluateSessionAction {
	_init_.Initialize()

	j := jsiiProxy_DomainEvents_VoiceIdEvaluateSessionAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDomainEvents_VoiceIdEvaluateSessionAction_Override(d DomainEvents_VoiceIdEvaluateSessionAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdEvaluateSessionAction",
		nil, // no parameters
		d,
	)
}

