package previewawseventsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogs",
		reflect.TypeOf((*CfnEventBusErrorLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnEventBusErrorLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsDestProps",
		reflect.TypeOf((*CfnEventBusErrorLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsFirehoseProps",
		reflect.TypeOf((*CfnEventBusErrorLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsLogGroupProps",
		reflect.TypeOf((*CfnEventBusErrorLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsOutputFormat",
		reflect.TypeOf((*CfnEventBusErrorLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnEventBusErrorLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnEventBusErrorLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnEventBusErrorLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnEventBusErrorLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnEventBusErrorLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnEventBusErrorLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnEventBusErrorLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnEventBusErrorLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsOutputFormat.S3",
		reflect.TypeOf((*CfnEventBusErrorLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnEventBusErrorLogsOutputFormat_S3_JSON,
			"PLAIN": CfnEventBusErrorLogsOutputFormat_S3_PLAIN,
			"W3C": CfnEventBusErrorLogsOutputFormat_S3_W3C,
			"PARQUET": CfnEventBusErrorLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsRecordFields",
		reflect.TypeOf((*CfnEventBusErrorLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnEventBusErrorLogsRecordFields_RESOURCE_ARN,
			"MESSAGE_TIMESTAMP_MS": CfnEventBusErrorLogsRecordFields_MESSAGE_TIMESTAMP_MS,
			"EVENT_BUS_NAME": CfnEventBusErrorLogsRecordFields_EVENT_BUS_NAME,
			"REQUEST_ID": CfnEventBusErrorLogsRecordFields_REQUEST_ID,
			"EVENT_ID": CfnEventBusErrorLogsRecordFields_EVENT_ID,
			"INVOCATION_ID": CfnEventBusErrorLogsRecordFields_INVOCATION_ID,
			"MESSAGE_TYPE": CfnEventBusErrorLogsRecordFields_MESSAGE_TYPE,
			"LOG_LEVEL": CfnEventBusErrorLogsRecordFields_LOG_LEVEL,
			"DETAILS": CfnEventBusErrorLogsRecordFields_DETAILS,
			"ERROR": CfnEventBusErrorLogsRecordFields_ERROR,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsS3Props",
		reflect.TypeOf((*CfnEventBusErrorLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogs",
		reflect.TypeOf((*CfnEventBusInfoLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnEventBusInfoLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsDestProps",
		reflect.TypeOf((*CfnEventBusInfoLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsFirehoseProps",
		reflect.TypeOf((*CfnEventBusInfoLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsLogGroupProps",
		reflect.TypeOf((*CfnEventBusInfoLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsOutputFormat",
		reflect.TypeOf((*CfnEventBusInfoLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnEventBusInfoLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnEventBusInfoLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnEventBusInfoLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnEventBusInfoLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnEventBusInfoLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnEventBusInfoLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnEventBusInfoLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnEventBusInfoLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsOutputFormat.S3",
		reflect.TypeOf((*CfnEventBusInfoLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnEventBusInfoLogsOutputFormat_S3_JSON,
			"PLAIN": CfnEventBusInfoLogsOutputFormat_S3_PLAIN,
			"W3C": CfnEventBusInfoLogsOutputFormat_S3_W3C,
			"PARQUET": CfnEventBusInfoLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsRecordFields",
		reflect.TypeOf((*CfnEventBusInfoLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnEventBusInfoLogsRecordFields_RESOURCE_ARN,
			"MESSAGE_TIMESTAMP_MS": CfnEventBusInfoLogsRecordFields_MESSAGE_TIMESTAMP_MS,
			"EVENT_BUS_NAME": CfnEventBusInfoLogsRecordFields_EVENT_BUS_NAME,
			"REQUEST_ID": CfnEventBusInfoLogsRecordFields_REQUEST_ID,
			"EVENT_ID": CfnEventBusInfoLogsRecordFields_EVENT_ID,
			"INVOCATION_ID": CfnEventBusInfoLogsRecordFields_INVOCATION_ID,
			"MESSAGE_TYPE": CfnEventBusInfoLogsRecordFields_MESSAGE_TYPE,
			"LOG_LEVEL": CfnEventBusInfoLogsRecordFields_LOG_LEVEL,
			"DETAILS": CfnEventBusInfoLogsRecordFields_DETAILS,
			"ERROR": CfnEventBusInfoLogsRecordFields_ERROR,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsS3Props",
		reflect.TypeOf((*CfnEventBusInfoLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusLogsMixin",
		reflect.TypeOf((*CfnEventBusLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventBusLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogs",
		reflect.TypeOf((*CfnEventBusTraceLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnEventBusTraceLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsDestProps",
		reflect.TypeOf((*CfnEventBusTraceLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsFirehoseProps",
		reflect.TypeOf((*CfnEventBusTraceLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsLogGroupProps",
		reflect.TypeOf((*CfnEventBusTraceLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsOutputFormat",
		reflect.TypeOf((*CfnEventBusTraceLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnEventBusTraceLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnEventBusTraceLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnEventBusTraceLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnEventBusTraceLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnEventBusTraceLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnEventBusTraceLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnEventBusTraceLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnEventBusTraceLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsOutputFormat.S3",
		reflect.TypeOf((*CfnEventBusTraceLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnEventBusTraceLogsOutputFormat_S3_JSON,
			"PLAIN": CfnEventBusTraceLogsOutputFormat_S3_PLAIN,
			"W3C": CfnEventBusTraceLogsOutputFormat_S3_W3C,
			"PARQUET": CfnEventBusTraceLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsRecordFields",
		reflect.TypeOf((*CfnEventBusTraceLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnEventBusTraceLogsRecordFields_RESOURCE_ARN,
			"MESSAGE_TIMESTAMP_MS": CfnEventBusTraceLogsRecordFields_MESSAGE_TIMESTAMP_MS,
			"EVENT_BUS_NAME": CfnEventBusTraceLogsRecordFields_EVENT_BUS_NAME,
			"REQUEST_ID": CfnEventBusTraceLogsRecordFields_REQUEST_ID,
			"EVENT_ID": CfnEventBusTraceLogsRecordFields_EVENT_ID,
			"INVOCATION_ID": CfnEventBusTraceLogsRecordFields_INVOCATION_ID,
			"MESSAGE_TYPE": CfnEventBusTraceLogsRecordFields_MESSAGE_TYPE,
			"LOG_LEVEL": CfnEventBusTraceLogsRecordFields_LOG_LEVEL,
			"DETAILS": CfnEventBusTraceLogsRecordFields_DETAILS,
			"ERROR": CfnEventBusTraceLogsRecordFields_ERROR,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsS3Props",
		reflect.TypeOf((*CfnEventBusTraceLogsS3Props)(nil)).Elem(),
	)
}
