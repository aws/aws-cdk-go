package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.voiceid@VoiceIdStartSessionAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdStartSessionAction := awscdkmixinspreview.Events.NewVoiceIdStartSessionAction()
//
// Experimental.
type VoiceIdStartSessionAction interface {
}

// The jsii proxy struct for VoiceIdStartSessionAction
type jsiiProxy_VoiceIdStartSessionAction struct {
	_ byte // padding
}

// Experimental.
func NewVoiceIdStartSessionAction() VoiceIdStartSessionAction {
	_init_.Initialize()

	j := jsiiProxy_VoiceIdStartSessionAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVoiceIdStartSessionAction_Override(v VoiceIdStartSessionAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction",
		nil, // no parameters
		v,
	)
}

// EventBridge event pattern for VoiceId Start Session Action.
// Experimental.
func VoiceIdStartSessionAction_VoiceIdStartSessionActionPattern(options *VoiceIdStartSessionAction_VoiceIdStartSessionActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateVoiceIdStartSessionAction_VoiceIdStartSessionActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdStartSessionAction",
		"voiceIdStartSessionActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

