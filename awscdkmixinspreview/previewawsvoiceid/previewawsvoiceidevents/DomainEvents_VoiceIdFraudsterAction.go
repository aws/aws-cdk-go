package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.voiceid@VoiceIdFraudsterAction event types for Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdFraudsterAction := #error#.NewVoiceIdFraudsterAction()
//
// Experimental.
type DomainEvents_VoiceIdFraudsterAction interface {
}

// The jsii proxy struct for DomainEvents_VoiceIdFraudsterAction
type jsiiProxy_DomainEvents_VoiceIdFraudsterAction struct {
	_ byte // padding
}

// Experimental.
func NewDomainEvents_VoiceIdFraudsterAction() DomainEvents_VoiceIdFraudsterAction {
	_init_.Initialize()

	j := jsiiProxy_DomainEvents_VoiceIdFraudsterAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdFraudsterAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDomainEvents_VoiceIdFraudsterAction_Override(d DomainEvents_VoiceIdFraudsterAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdFraudsterAction",
		nil, // no parameters
		d,
	)
}

