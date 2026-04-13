package previewawsivsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ivs@RecordingStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recordingStateChange := awscdkmixinspreview.Events.NewRecordingStateChange()
//
// Experimental.
type RecordingStateChange interface {
}

// The jsii proxy struct for RecordingStateChange
type jsiiProxy_RecordingStateChange struct {
	_ byte // padding
}

// Experimental.
func NewRecordingStateChange() RecordingStateChange {
	_init_.Initialize()

	j := jsiiProxy_RecordingStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.events.RecordingStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRecordingStateChange_Override(r RecordingStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.events.RecordingStateChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for IVS Recording State Change.
// Experimental.
func RecordingStateChange_EventPattern(options *RecordingStateChange_RecordingStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRecordingStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ivs.events.RecordingStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

