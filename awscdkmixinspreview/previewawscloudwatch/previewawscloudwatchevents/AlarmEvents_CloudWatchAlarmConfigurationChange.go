package previewawscloudwatchevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.cloudwatch@CloudWatchAlarmConfigurationChange event types for Alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchAlarmConfigurationChange := #error#.NewCloudWatchAlarmConfigurationChange()
//
// Experimental.
type AlarmEvents_CloudWatchAlarmConfigurationChange interface {
}

// The jsii proxy struct for AlarmEvents_CloudWatchAlarmConfigurationChange
type jsiiProxy_AlarmEvents_CloudWatchAlarmConfigurationChange struct {
	_ byte // padding
}

// Experimental.
func NewAlarmEvents_CloudWatchAlarmConfigurationChange() AlarmEvents_CloudWatchAlarmConfigurationChange {
	_init_.Initialize()

	j := jsiiProxy_AlarmEvents_CloudWatchAlarmConfigurationChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmConfigurationChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAlarmEvents_CloudWatchAlarmConfigurationChange_Override(a AlarmEvents_CloudWatchAlarmConfigurationChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmConfigurationChange",
		nil, // no parameters
		a,
	)
}

