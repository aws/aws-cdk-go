package previewawscloudwatchevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.cloudwatch@CloudWatchAlarmStateChange event types for Alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchAlarmStateChange := #error#.NewCloudWatchAlarmStateChange()
//
// Experimental.
type AlarmEvents_CloudWatchAlarmStateChange interface {
}

// The jsii proxy struct for AlarmEvents_CloudWatchAlarmStateChange
type jsiiProxy_AlarmEvents_CloudWatchAlarmStateChange struct {
	_ byte // padding
}

// Experimental.
func NewAlarmEvents_CloudWatchAlarmStateChange() AlarmEvents_CloudWatchAlarmStateChange {
	_init_.Initialize()

	j := jsiiProxy_AlarmEvents_CloudWatchAlarmStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAlarmEvents_CloudWatchAlarmStateChange_Override(a AlarmEvents_CloudWatchAlarmStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmStateChange",
		nil, // no parameters
		a,
	)
}

