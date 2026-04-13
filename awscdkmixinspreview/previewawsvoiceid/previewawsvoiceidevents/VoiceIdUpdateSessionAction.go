package previewawsvoiceidevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.voiceid@VoiceIdUpdateSessionAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdUpdateSessionAction := awscdkmixinspreview.Events.NewVoiceIdUpdateSessionAction()
//
// Experimental.
type VoiceIdUpdateSessionAction interface {
}

// The jsii proxy struct for VoiceIdUpdateSessionAction
type jsiiProxy_VoiceIdUpdateSessionAction struct {
	_ byte // padding
}

// Experimental.
func NewVoiceIdUpdateSessionAction() VoiceIdUpdateSessionAction {
	_init_.Initialize()

	j := jsiiProxy_VoiceIdUpdateSessionAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdUpdateSessionAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVoiceIdUpdateSessionAction_Override(v VoiceIdUpdateSessionAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdUpdateSessionAction",
		nil, // no parameters
		v,
	)
}

// EventBridge event pattern for VoiceId Update Session Action.
// Experimental.
func VoiceIdUpdateSessionAction_EventPattern(options *VoiceIdUpdateSessionAction_VoiceIdUpdateSessionActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateVoiceIdUpdateSessionAction_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_voiceid.events.VoiceIdUpdateSessionAction",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

