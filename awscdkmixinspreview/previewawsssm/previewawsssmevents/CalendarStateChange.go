package previewawsssmevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ssm@CalendarStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   calendarStateChange := awscdkmixinspreview.Events.NewCalendarStateChange()
//
// Experimental.
type CalendarStateChange interface {
}

// The jsii proxy struct for CalendarStateChange
type jsiiProxy_CalendarStateChange struct {
	_ byte // padding
}

// Experimental.
func NewCalendarStateChange() CalendarStateChange {
	_init_.Initialize()

	j := jsiiProxy_CalendarStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.events.CalendarStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCalendarStateChange_Override(c CalendarStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.events.CalendarStateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Calendar State Change.
// Experimental.
func CalendarStateChange_EventPattern(options *CalendarStateChange_CalendarStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCalendarStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssm.events.CalendarStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

