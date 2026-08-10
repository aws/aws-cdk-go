package previewawstranscribeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transcribe@TranscribeJobStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transcribeJobStateChange := awscdkmixinspreview.Events.NewTranscribeJobStateChange()
//
// Experimental.
type TranscribeJobStateChange interface {
}

// The jsii proxy struct for TranscribeJobStateChange
type jsiiProxy_TranscribeJobStateChange struct {
	_ byte // padding
}

// Experimental.
func NewTranscribeJobStateChange() TranscribeJobStateChange {
	_init_.Initialize()

	j := jsiiProxy_TranscribeJobStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transcribe.events.TranscribeJobStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTranscribeJobStateChange_Override(t TranscribeJobStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transcribe.events.TranscribeJobStateChange",
		nil, // no parameters
		t,
	)
}

// EventBridge event pattern for Transcribe Job State Change.
// Experimental.
func TranscribeJobStateChange_EventPattern(options *TranscribeJobStateChange_TranscribeJobStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateTranscribeJobStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transcribe.events.TranscribeJobStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

