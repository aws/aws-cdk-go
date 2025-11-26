package previewawscloudwatchevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
)

// EventBridge event patterns for Alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var alarmRef IAlarmRef
//
//   alarmEvents := awscdkmixinspreview.Events.AlarmEvents_FromAlarm(alarmRef)
//
// Experimental.
type AlarmEvents interface {
	// EventBridge event pattern for Alarm CloudWatch Alarm Configuration Change.
	// Experimental.
	CloudWatchAlarmConfigurationChangePattern(options *AlarmEvents_CloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangeProps) *awsevents.EventPattern
	// EventBridge event pattern for Alarm CloudWatch Alarm State Change.
	// Experimental.
	CloudWatchAlarmStateChangePattern(options *AlarmEvents_CloudWatchAlarmStateChange_CloudWatchAlarmStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for AlarmEvents
type jsiiProxy_AlarmEvents struct {
	_ byte // padding
}

// Create AlarmEvents from a Alarm reference.
// Experimental.
func AlarmEvents_FromAlarm(alarmRef interfacesawscloudwatch.IAlarmRef) AlarmEvents {
	_init_.Initialize()

	if err := validateAlarmEvents_FromAlarmParameters(alarmRef); err != nil {
		panic(err)
	}
	var returns AlarmEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents",
		"fromAlarm",
		[]interface{}{alarmRef},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlarmEvents) CloudWatchAlarmConfigurationChangePattern(options *AlarmEvents_CloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangeProps) *awsevents.EventPattern {
	if err := a.validateCloudWatchAlarmConfigurationChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"cloudWatchAlarmConfigurationChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlarmEvents) CloudWatchAlarmStateChangePattern(options *AlarmEvents_CloudWatchAlarmStateChange_CloudWatchAlarmStateChangeProps) *awsevents.EventPattern {
	if err := a.validateCloudWatchAlarmStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"cloudWatchAlarmStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

