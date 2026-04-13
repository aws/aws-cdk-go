package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.voiceid@VoiceIdFraudsterAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdFraudsterAction := awscdkmixinspreview.Events.NewVoiceIdFraudsterAction()
//
// Experimental.
type VoiceIdFraudsterAction interface {
}

// The jsii proxy struct for VoiceIdFraudsterAction
type jsiiProxy_VoiceIdFraudsterAction struct {
	_ byte // padding
}

// Experimental.
func NewVoiceIdFraudsterAction() VoiceIdFraudsterAction {
	_init_.Initialize()

	j := jsiiProxy_VoiceIdFraudsterAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdFraudsterAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVoiceIdFraudsterAction_Override(v VoiceIdFraudsterAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdFraudsterAction",
		nil, // no parameters
		v,
	)
}

// EventBridge event pattern for VoiceId Fraudster Action.
// Experimental.
func VoiceIdFraudsterAction_EventPattern(options *VoiceIdFraudsterAction_VoiceIdFraudsterActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateVoiceIdFraudsterAction_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdFraudsterAction",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

