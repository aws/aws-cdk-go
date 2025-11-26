package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.voiceid@VoiceIdStartSessionAction event types for Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdStartSessionAction := #error#.NewVoiceIdStartSessionAction()
//
// Experimental.
type DomainEvents_VoiceIdStartSessionAction interface {
}

// The jsii proxy struct for DomainEvents_VoiceIdStartSessionAction
type jsiiProxy_DomainEvents_VoiceIdStartSessionAction struct {
	_ byte // padding
}

// Experimental.
func NewDomainEvents_VoiceIdStartSessionAction() DomainEvents_VoiceIdStartSessionAction {
	_init_.Initialize()

	j := jsiiProxy_DomainEvents_VoiceIdStartSessionAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDomainEvents_VoiceIdStartSessionAction_Override(d DomainEvents_VoiceIdStartSessionAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdStartSessionAction",
		nil, // no parameters
		d,
	)
}

