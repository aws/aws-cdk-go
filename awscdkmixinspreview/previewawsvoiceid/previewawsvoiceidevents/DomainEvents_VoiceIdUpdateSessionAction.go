package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.voiceid@VoiceIdUpdateSessionAction event types for Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdUpdateSessionAction := #error#.NewVoiceIdUpdateSessionAction()
//
// Experimental.
type DomainEvents_VoiceIdUpdateSessionAction interface {
}

// The jsii proxy struct for DomainEvents_VoiceIdUpdateSessionAction
type jsiiProxy_DomainEvents_VoiceIdUpdateSessionAction struct {
	_ byte // padding
}

// Experimental.
func NewDomainEvents_VoiceIdUpdateSessionAction() DomainEvents_VoiceIdUpdateSessionAction {
	_init_.Initialize()

	j := jsiiProxy_DomainEvents_VoiceIdUpdateSessionAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdUpdateSessionAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDomainEvents_VoiceIdUpdateSessionAction_Override(d DomainEvents_VoiceIdUpdateSessionAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdUpdateSessionAction",
		nil, // no parameters
		d,
	)
}

