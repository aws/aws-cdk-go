package previewawspcsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterLogsMixin",
		reflect.TypeOf((*CfnClusterLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogs",
		reflect.TypeOf((*CfnClusterPcsJobcompLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterPcsJobcompLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsDestProps",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsOutputFormat",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterPcsJobcompLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterPcsJobcompLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnClusterPcsJobcompLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnClusterPcsJobcompLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterPcsJobcompLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnClusterPcsJobcompLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterPcsJobcompLogsOutputFormat_S3_JSON,
			"PLAIN": CfnClusterPcsJobcompLogsOutputFormat_S3_PLAIN,
			"W3C": CfnClusterPcsJobcompLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterPcsJobcompLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsRecordFields",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ID": CfnClusterPcsJobcompLogsRecordFields_RESOURCE_ID,
			"RESOURCE_TYPE": CfnClusterPcsJobcompLogsRecordFields_RESOURCE_TYPE,
			"EVENT_TIMESTAMP": CfnClusterPcsJobcompLogsRecordFields_EVENT_TIMESTAMP,
			"SCHEDULER_TYPE": CfnClusterPcsJobcompLogsRecordFields_SCHEDULER_TYPE,
			"SCHEDULER_MAJOR_VERSION": CfnClusterPcsJobcompLogsRecordFields_SCHEDULER_MAJOR_VERSION,
			"FIELDS": CfnClusterPcsJobcompLogsRecordFields_FIELDS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsS3Props",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogs",
		reflect.TypeOf((*CfnClusterPcsSchedulerAuditLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterPcsSchedulerAuditLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsDestProps",
		reflect.TypeOf((*CfnClusterPcsSchedulerAuditLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterPcsSchedulerAuditLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterPcsSchedulerAuditLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsOutputFormat",
		reflect.TypeOf((*CfnClusterPcsSchedulerAuditLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterPcsSchedulerAuditLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterPcsSchedulerAuditLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterPcsSchedulerAuditLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnClusterPcsSchedulerAuditLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnClusterPcsSchedulerAuditLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterPcsSchedulerAuditLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterPcsSchedulerAuditLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnClusterPcsSchedulerAuditLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterPcsSchedulerAuditLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterPcsSchedulerAuditLogsOutputFormat_S3_JSON,
			"PLAIN": CfnClusterPcsSchedulerAuditLogsOutputFormat_S3_PLAIN,
			"W3C": CfnClusterPcsSchedulerAuditLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterPcsSchedulerAuditLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsRecordFields",
		reflect.TypeOf((*CfnClusterPcsSchedulerAuditLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ID": CfnClusterPcsSchedulerAuditLogsRecordFields_RESOURCE_ID,
			"RESOURCE_TYPE": CfnClusterPcsSchedulerAuditLogsRecordFields_RESOURCE_TYPE,
			"EVENT_TIMESTAMP": CfnClusterPcsSchedulerAuditLogsRecordFields_EVENT_TIMESTAMP,
			"LOG_LEVEL": CfnClusterPcsSchedulerAuditLogsRecordFields_LOG_LEVEL,
			"LOG_NAME": CfnClusterPcsSchedulerAuditLogsRecordFields_LOG_NAME,
			"SCHEDULER_TYPE": CfnClusterPcsSchedulerAuditLogsRecordFields_SCHEDULER_TYPE,
			"SCHEDULER_MAJOR_VERSION": CfnClusterPcsSchedulerAuditLogsRecordFields_SCHEDULER_MAJOR_VERSION,
			"SCHEDULER_PATCH_VERSION": CfnClusterPcsSchedulerAuditLogsRecordFields_SCHEDULER_PATCH_VERSION,
			"NODE_TYPE": CfnClusterPcsSchedulerAuditLogsRecordFields_NODE_TYPE,
			"LOG_TYPE": CfnClusterPcsSchedulerAuditLogsRecordFields_LOG_TYPE,
			"MESSAGE": CfnClusterPcsSchedulerAuditLogsRecordFields_MESSAGE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsS3Props",
		reflect.TypeOf((*CfnClusterPcsSchedulerAuditLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogs",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterPcsSchedulerLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsDestProps",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsOutputFormat",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterPcsSchedulerLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterPcsSchedulerLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnClusterPcsSchedulerLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnClusterPcsSchedulerLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterPcsSchedulerLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnClusterPcsSchedulerLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterPcsSchedulerLogsOutputFormat_S3_JSON,
			"PLAIN": CfnClusterPcsSchedulerLogsOutputFormat_S3_PLAIN,
			"W3C": CfnClusterPcsSchedulerLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterPcsSchedulerLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsRecordFields",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ID": CfnClusterPcsSchedulerLogsRecordFields_RESOURCE_ID,
			"RESOURCE_TYPE": CfnClusterPcsSchedulerLogsRecordFields_RESOURCE_TYPE,
			"EVENT_TIMESTAMP": CfnClusterPcsSchedulerLogsRecordFields_EVENT_TIMESTAMP,
			"LOG_LEVEL": CfnClusterPcsSchedulerLogsRecordFields_LOG_LEVEL,
			"LOG_NAME": CfnClusterPcsSchedulerLogsRecordFields_LOG_NAME,
			"SCHEDULER_TYPE": CfnClusterPcsSchedulerLogsRecordFields_SCHEDULER_TYPE,
			"SCHEDULER_MAJOR_VERSION": CfnClusterPcsSchedulerLogsRecordFields_SCHEDULER_MAJOR_VERSION,
			"SCHEDULER_PATCH_VERSION": CfnClusterPcsSchedulerLogsRecordFields_SCHEDULER_PATCH_VERSION,
			"NODE_TYPE": CfnClusterPcsSchedulerLogsRecordFields_NODE_TYPE,
			"MESSAGE": CfnClusterPcsSchedulerLogsRecordFields_MESSAGE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsS3Props",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsS3Props)(nil)).Elem(),
	)
}
