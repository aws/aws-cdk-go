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
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmConfigurationChange",
		reflect.TypeOf((*CloudWatchAlarmConfigurationChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CloudWatchAlarmConfigurationChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmConfigurationChange.CloudWatchAlarmConfigurationChangeProps",
		reflect.TypeOf((*CloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmConfigurationChange.Configuration",
		reflect.TypeOf((*CloudWatchAlarmConfigurationChange_Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmConfigurationChange.ConfigurationItem",
		reflect.TypeOf((*CloudWatchAlarmConfigurationChange_ConfigurationItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmConfigurationChange.Metric",
		reflect.TypeOf((*CloudWatchAlarmConfigurationChange_Metric)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmConfigurationChange.MetricStat",
		reflect.TypeOf((*CloudWatchAlarmConfigurationChange_MetricStat)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmConfigurationChange.State",
		reflect.TypeOf((*CloudWatchAlarmConfigurationChange_State)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmStateChange",
		reflect.TypeOf((*CloudWatchAlarmStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CloudWatchAlarmStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmStateChange.CloudWatchAlarmStateChangeProps",
		reflect.TypeOf((*CloudWatchAlarmStateChange_CloudWatchAlarmStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmStateChange.Configuration",
		reflect.TypeOf((*CloudWatchAlarmStateChange_Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmStateChange.ConfigurationItem",
		reflect.TypeOf((*CloudWatchAlarmStateChange_ConfigurationItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmStateChange.Metric",
		reflect.TypeOf((*CloudWatchAlarmStateChange_Metric)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmStateChange.MetricStat",
		reflect.TypeOf((*CloudWatchAlarmStateChange_MetricStat)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmStateChange.State",
		reflect.TypeOf((*CloudWatchAlarmStateChange_State)(nil)).Elem(),
	)
}
