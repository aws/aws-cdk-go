package previewawseventsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.events@ScheduledJson.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scheduledJson := awscdkmixinspreview.Events.NewScheduledJson()
//
// Experimental.
type ScheduledJson interface {
}

// The jsii proxy struct for ScheduledJson
type jsiiProxy_ScheduledJson struct {
	_ byte // padding
}

// Experimental.
func NewScheduledJson() ScheduledJson {
	_init_.Initialize()

	j := jsiiProxy_ScheduledJson{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.events.ScheduledJson",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewScheduledJson_Override(s ScheduledJson) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.events.ScheduledJson",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Scheduled Event.
// Experimental.
func ScheduledJson_ScheduledJsonPattern(options *ScheduledJson_ScheduledJsonProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateScheduledJson_ScheduledJsonPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_events.events.ScheduledJson",
		"scheduledJsonPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

