package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.voiceid@VoiceIdEvaluateSessionAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdEvaluateSessionAction := awscdkmixinspreview.Events.NewVoiceIdEvaluateSessionAction()
//
// Experimental.
type VoiceIdEvaluateSessionAction interface {
}

// The jsii proxy struct for VoiceIdEvaluateSessionAction
type jsiiProxy_VoiceIdEvaluateSessionAction struct {
	_ byte // padding
}

// Experimental.
func NewVoiceIdEvaluateSessionAction() VoiceIdEvaluateSessionAction {
	_init_.Initialize()

	j := jsiiProxy_VoiceIdEvaluateSessionAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVoiceIdEvaluateSessionAction_Override(v VoiceIdEvaluateSessionAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction",
		nil, // no parameters
		v,
	)
}

// EventBridge event pattern for VoiceId Evaluate Session Action.
// Experimental.
func VoiceIdEvaluateSessionAction_VoiceIdEvaluateSessionActionPattern(options *VoiceIdEvaluateSessionAction_VoiceIdEvaluateSessionActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateVoiceIdEvaluateSessionAction_VoiceIdEvaluateSessionActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdEvaluateSessionAction",
		"voiceIdEvaluateSessionActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

