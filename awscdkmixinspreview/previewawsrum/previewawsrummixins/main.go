package previewawsrummixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorLogsMixin",
		reflect.TypeOf((*CfnAppMonitorLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppMonitorLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorMixinProps",
		reflect.TypeOf((*CfnAppMonitorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin",
		reflect.TypeOf((*CfnAppMonitorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppMonitorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.AppMonitorConfigurationProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_AppMonitorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.CustomEventsProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_CustomEventsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.DeobfuscationConfigurationProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_DeobfuscationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.JavaScriptSourceMapsProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_JavaScriptSourceMapsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.MetricDefinitionProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_MetricDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.MetricDestinationProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_MetricDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.ResourcePolicyProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_ResourcePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogs",
		reflect.TypeOf((*CfnAppMonitorRumOtelLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnAppMonitorRumOtelLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogsFirehoseProps",
		reflect.TypeOf((*CfnAppMonitorRumOtelLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogsLogGroupProps",
		reflect.TypeOf((*CfnAppMonitorRumOtelLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogsOutputFormat",
		reflect.TypeOf((*CfnAppMonitorRumOtelLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnAppMonitorRumOtelLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnAppMonitorRumOtelLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAppMonitorRumOtelLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnAppMonitorRumOtelLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnAppMonitorRumOtelLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnAppMonitorRumOtelLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAppMonitorRumOtelLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnAppMonitorRumOtelLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogsOutputFormat.S3",
		reflect.TypeOf((*CfnAppMonitorRumOtelLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAppMonitorRumOtelLogsOutputFormat_S3_JSON,
			"PLAIN": CfnAppMonitorRumOtelLogsOutputFormat_S3_PLAIN,
			"W3C": CfnAppMonitorRumOtelLogsOutputFormat_S3_W3C,
			"PARQUET": CfnAppMonitorRumOtelLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogsS3Props",
		reflect.TypeOf((*CfnAppMonitorRumOtelLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelSpans",
		reflect.TypeOf((*CfnAppMonitorRumOtelSpans)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnAppMonitorRumOtelSpans{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelSpansFirehoseProps",
		reflect.TypeOf((*CfnAppMonitorRumOtelSpansFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelSpansLogGroupProps",
		reflect.TypeOf((*CfnAppMonitorRumOtelSpansLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelSpansOutputFormat",
		reflect.TypeOf((*CfnAppMonitorRumOtelSpansOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnAppMonitorRumOtelSpansOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelSpansOutputFormat.Firehose",
		reflect.TypeOf((*CfnAppMonitorRumOtelSpansOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAppMonitorRumOtelSpansOutputFormat_Firehose_JSON,
			"PLAIN": CfnAppMonitorRumOtelSpansOutputFormat_Firehose_PLAIN,
			"RAW": CfnAppMonitorRumOtelSpansOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelSpansOutputFormat.LogGroup",
		reflect.TypeOf((*CfnAppMonitorRumOtelSpansOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAppMonitorRumOtelSpansOutputFormat_LogGroup_PLAIN,
			"JSON": CfnAppMonitorRumOtelSpansOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelSpansOutputFormat.S3",
		reflect.TypeOf((*CfnAppMonitorRumOtelSpansOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAppMonitorRumOtelSpansOutputFormat_S3_JSON,
			"PLAIN": CfnAppMonitorRumOtelSpansOutputFormat_S3_PLAIN,
			"W3C": CfnAppMonitorRumOtelSpansOutputFormat_S3_W3C,
			"PARQUET": CfnAppMonitorRumOtelSpansOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelSpansS3Props",
		reflect.TypeOf((*CfnAppMonitorRumOtelSpansS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumTelemetryLogs",
		reflect.TypeOf((*CfnAppMonitorRumTelemetryLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnAppMonitorRumTelemetryLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumTelemetryLogsFirehoseProps",
		reflect.TypeOf((*CfnAppMonitorRumTelemetryLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumTelemetryLogsLogGroupProps",
		reflect.TypeOf((*CfnAppMonitorRumTelemetryLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumTelemetryLogsOutputFormat",
		reflect.TypeOf((*CfnAppMonitorRumTelemetryLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnAppMonitorRumTelemetryLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumTelemetryLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnAppMonitorRumTelemetryLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAppMonitorRumTelemetryLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnAppMonitorRumTelemetryLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnAppMonitorRumTelemetryLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumTelemetryLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnAppMonitorRumTelemetryLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAppMonitorRumTelemetryLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnAppMonitorRumTelemetryLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumTelemetryLogsOutputFormat.S3",
		reflect.TypeOf((*CfnAppMonitorRumTelemetryLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAppMonitorRumTelemetryLogsOutputFormat_S3_JSON,
			"PLAIN": CfnAppMonitorRumTelemetryLogsOutputFormat_S3_PLAIN,
			"W3C": CfnAppMonitorRumTelemetryLogsOutputFormat_S3_W3C,
			"PARQUET": CfnAppMonitorRumTelemetryLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumTelemetryLogsS3Props",
		reflect.TypeOf((*CfnAppMonitorRumTelemetryLogsS3Props)(nil)).Elem(),
	)
}
