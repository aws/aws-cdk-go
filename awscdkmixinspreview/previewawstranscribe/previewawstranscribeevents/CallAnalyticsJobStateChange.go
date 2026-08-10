package previewawstranscribeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transcribe@CallAnalyticsJobStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   callAnalyticsJobStateChange := awscdkmixinspreview.Events.NewCallAnalyticsJobStateChange()
//
// Experimental.
type CallAnalyticsJobStateChange interface {
}

// The jsii proxy struct for CallAnalyticsJobStateChange
type jsiiProxy_CallAnalyticsJobStateChange struct {
	_ byte // padding
}

// Experimental.
func NewCallAnalyticsJobStateChange() CallAnalyticsJobStateChange {
	_init_.Initialize()

	j := jsiiProxy_CallAnalyticsJobStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transcribe.events.CallAnalyticsJobStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCallAnalyticsJobStateChange_Override(c CallAnalyticsJobStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transcribe.events.CallAnalyticsJobStateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Call Analytics Job State Change.
// Experimental.
func CallAnalyticsJobStateChange_EventPattern(options *CallAnalyticsJobStateChange_CallAnalyticsJobStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCallAnalyticsJobStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transcribe.events.CallAnalyticsJobStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

