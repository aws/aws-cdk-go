package previewawslogsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogs",
		reflect.TypeOf((*CfnLogGroupAuditLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnLogGroupAuditLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsDestProps",
		reflect.TypeOf((*CfnLogGroupAuditLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsFirehoseProps",
		reflect.TypeOf((*CfnLogGroupAuditLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsLogGroupProps",
		reflect.TypeOf((*CfnLogGroupAuditLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsOutputFormat",
		reflect.TypeOf((*CfnLogGroupAuditLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnLogGroupAuditLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnLogGroupAuditLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLogGroupAuditLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnLogGroupAuditLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnLogGroupAuditLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnLogGroupAuditLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnLogGroupAuditLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnLogGroupAuditLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsOutputFormat.S3",
		reflect.TypeOf((*CfnLogGroupAuditLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLogGroupAuditLogsOutputFormat_S3_JSON,
			"PLAIN": CfnLogGroupAuditLogsOutputFormat_S3_PLAIN,
			"W3C": CfnLogGroupAuditLogsOutputFormat_S3_W3C,
			"PARQUET": CfnLogGroupAuditLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsRecordFields",
		reflect.TypeOf((*CfnLogGroupAuditLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"POLICYNAME": CfnLogGroupAuditLogsRecordFields_POLICYNAME,
			"CLOUDWATCHLOGS_EVENTTIMESTAMP": CfnLogGroupAuditLogsRecordFields_CLOUDWATCHLOGS_EVENTTIMESTAMP,
			"CLOUDWATCHLOGS_LOGSTREAM": CfnLogGroupAuditLogsRecordFields_CLOUDWATCHLOGS_LOGSTREAM,
			"AUDITTIMESTAMP": CfnLogGroupAuditLogsRecordFields_AUDITTIMESTAMP,
			"RESOURCEARN": CfnLogGroupAuditLogsRecordFields_RESOURCEARN,
			"RESOURCENAME": CfnLogGroupAuditLogsRecordFields_RESOURCENAME,
			"DATAIDENTIFIERS": CfnLogGroupAuditLogsRecordFields_DATAIDENTIFIERS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsS3Props",
		reflect.TypeOf((*CfnLogGroupAuditLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupLogsMixin",
		reflect.TypeOf((*CfnLogGroupLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogGroupLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
