package previewawscloudwatchevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents",
		reflect.TypeOf((*AlarmEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "cloudWatchAlarmConfigurationChangePattern", GoMethod: "CloudWatchAlarmConfigurationChangePattern"},
			_jsii_.MemberMethod{JsiiMethod: "cloudWatchAlarmStateChangePattern", GoMethod: "CloudWatchAlarmStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_AlarmEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmConfigurationChange",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmConfigurationChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AlarmEvents_CloudWatchAlarmConfigurationChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmConfigurationChange.CloudWatchAlarmConfigurationChangeProps",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmConfigurationChange.Configuration",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmConfigurationChange_Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmConfigurationChange.ConfigurationItem",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmConfigurationChange_ConfigurationItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmConfigurationChange.Metric",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmConfigurationChange_Metric)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmConfigurationChange.MetricStat",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmConfigurationChange_MetricStat)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmConfigurationChange.State",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmConfigurationChange_State)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmStateChange",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AlarmEvents_CloudWatchAlarmStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmStateChange.CloudWatchAlarmStateChangeProps",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmStateChange_CloudWatchAlarmStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmStateChange.Configuration",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmStateChange_Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmStateChange.ConfigurationItem",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmStateChange_ConfigurationItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmStateChange.Metric",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmStateChange_Metric)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmStateChange.MetricStat",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmStateChange_MetricStat)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.AlarmEvents.CloudWatchAlarmStateChange.State",
		reflect.TypeOf((*AlarmEvents_CloudWatchAlarmStateChange_State)(nil)).Elem(),
	)
}
