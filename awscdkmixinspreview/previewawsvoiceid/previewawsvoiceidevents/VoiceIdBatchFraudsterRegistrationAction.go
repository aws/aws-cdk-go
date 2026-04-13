package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.voiceid@VoiceIdBatchFraudsterRegistrationAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdBatchFraudsterRegistrationAction := awscdkmixinspreview.Events.NewVoiceIdBatchFraudsterRegistrationAction()
//
// Experimental.
type VoiceIdBatchFraudsterRegistrationAction interface {
}

// The jsii proxy struct for VoiceIdBatchFraudsterRegistrationAction
type jsiiProxy_VoiceIdBatchFraudsterRegistrationAction struct {
	_ byte // padding
}

// Experimental.
func NewVoiceIdBatchFraudsterRegistrationAction() VoiceIdBatchFraudsterRegistrationAction {
	_init_.Initialize()

	j := jsiiProxy_VoiceIdBatchFraudsterRegistrationAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchFraudsterRegistrationAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVoiceIdBatchFraudsterRegistrationAction_Override(v VoiceIdBatchFraudsterRegistrationAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchFraudsterRegistrationAction",
		nil, // no parameters
		v,
	)
}

// EventBridge event pattern for VoiceId Batch Fraudster Registration Action.
// Experimental.
func VoiceIdBatchFraudsterRegistrationAction_EventPattern(options *VoiceIdBatchFraudsterRegistrationAction_VoiceIdBatchFraudsterRegistrationActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateVoiceIdBatchFraudsterRegistrationAction_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdBatchFraudsterRegistrationAction",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

