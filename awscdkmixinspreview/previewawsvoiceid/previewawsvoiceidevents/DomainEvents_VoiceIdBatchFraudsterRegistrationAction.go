package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.voiceid@VoiceIdBatchFraudsterRegistrationAction event types for Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdBatchFraudsterRegistrationAction := #error#.NewVoiceIdBatchFraudsterRegistrationAction()
//
// Experimental.
type DomainEvents_VoiceIdBatchFraudsterRegistrationAction interface {
}

// The jsii proxy struct for DomainEvents_VoiceIdBatchFraudsterRegistrationAction
type jsiiProxy_DomainEvents_VoiceIdBatchFraudsterRegistrationAction struct {
	_ byte // padding
}

// Experimental.
func NewDomainEvents_VoiceIdBatchFraudsterRegistrationAction() DomainEvents_VoiceIdBatchFraudsterRegistrationAction {
	_init_.Initialize()

	j := jsiiProxy_DomainEvents_VoiceIdBatchFraudsterRegistrationAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchFraudsterRegistrationAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDomainEvents_VoiceIdBatchFraudsterRegistrationAction_Override(d DomainEvents_VoiceIdBatchFraudsterRegistrationAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.DomainEvents.VoiceIdBatchFraudsterRegistrationAction",
		nil, // no parameters
		d,
	)
}

