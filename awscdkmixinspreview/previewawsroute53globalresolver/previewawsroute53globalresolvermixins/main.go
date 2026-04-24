package previewawsroute53globalresolvermixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogs",
		reflect.TypeOf((*CfnGlobalResolverGlobalResolverLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnGlobalResolverGlobalResolverLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsDestProps",
		reflect.TypeOf((*CfnGlobalResolverGlobalResolverLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsFirehoseProps",
		reflect.TypeOf((*CfnGlobalResolverGlobalResolverLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsLogGroupProps",
		reflect.TypeOf((*CfnGlobalResolverGlobalResolverLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsOutputFormat",
		reflect.TypeOf((*CfnGlobalResolverGlobalResolverLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnGlobalResolverGlobalResolverLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnGlobalResolverGlobalResolverLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnGlobalResolverGlobalResolverLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnGlobalResolverGlobalResolverLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnGlobalResolverGlobalResolverLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnGlobalResolverGlobalResolverLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnGlobalResolverGlobalResolverLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnGlobalResolverGlobalResolverLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsOutputFormat.S3",
		reflect.TypeOf((*CfnGlobalResolverGlobalResolverLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnGlobalResolverGlobalResolverLogsOutputFormat_S3_JSON,
			"PLAIN": CfnGlobalResolverGlobalResolverLogsOutputFormat_S3_PLAIN,
			"W3C": CfnGlobalResolverGlobalResolverLogsOutputFormat_S3_W3C,
			"PARQUET": CfnGlobalResolverGlobalResolverLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsRecordFields",
		reflect.TypeOf((*CfnGlobalResolverGlobalResolverLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"ACTION_NAME": CfnGlobalResolverGlobalResolverLogsRecordFields_ACTION_NAME,
			"ACTIVITY_NAME": CfnGlobalResolverGlobalResolverLogsRecordFields_ACTIVITY_NAME,
			"ANSWERS": CfnGlobalResolverGlobalResolverLogsRecordFields_ANSWERS,
			"CATEGORY_NAME": CfnGlobalResolverGlobalResolverLogsRecordFields_CATEGORY_NAME,
			"CLASS_NAME": CfnGlobalResolverGlobalResolverLogsRecordFields_CLASS_NAME,
			"CONNECTION_INFO": CfnGlobalResolverGlobalResolverLogsRecordFields_CONNECTION_INFO,
			"DISPOSITION": CfnGlobalResolverGlobalResolverLogsRecordFields_DISPOSITION,
			"DISPOSITION_ID": CfnGlobalResolverGlobalResolverLogsRecordFields_DISPOSITION_ID,
			"DURATION": CfnGlobalResolverGlobalResolverLogsRecordFields_DURATION,
			"END_TIME": CfnGlobalResolverGlobalResolverLogsRecordFields_END_TIME,
			"ENRICHMENTS": CfnGlobalResolverGlobalResolverLogsRecordFields_ENRICHMENTS,
			"MESSAGE": CfnGlobalResolverGlobalResolverLogsRecordFields_MESSAGE,
			"QUERY": CfnGlobalResolverGlobalResolverLogsRecordFields_QUERY,
			"QUERY_TIME": CfnGlobalResolverGlobalResolverLogsRecordFields_QUERY_TIME,
			"RCODE": CfnGlobalResolverGlobalResolverLogsRecordFields_RCODE,
			"RCODE_ID": CfnGlobalResolverGlobalResolverLogsRecordFields_RCODE_ID,
			"RESPONSE_TIME": CfnGlobalResolverGlobalResolverLogsRecordFields_RESPONSE_TIME,
			"SEVERITY": CfnGlobalResolverGlobalResolverLogsRecordFields_SEVERITY,
			"START_TIME": CfnGlobalResolverGlobalResolverLogsRecordFields_START_TIME,
			"STATUS": CfnGlobalResolverGlobalResolverLogsRecordFields_STATUS,
			"STATUS_ID": CfnGlobalResolverGlobalResolverLogsRecordFields_STATUS_ID,
			"TYPE_NAME": CfnGlobalResolverGlobalResolverLogsRecordFields_TYPE_NAME,
			"ACTION_ID": CfnGlobalResolverGlobalResolverLogsRecordFields_ACTION_ID,
			"ACTIVITY_ID": CfnGlobalResolverGlobalResolverLogsRecordFields_ACTIVITY_ID,
			"CATEGORY_UID": CfnGlobalResolverGlobalResolverLogsRecordFields_CATEGORY_UID,
			"CLASS_UID": CfnGlobalResolverGlobalResolverLogsRecordFields_CLASS_UID,
			"CLOUD": CfnGlobalResolverGlobalResolverLogsRecordFields_CLOUD,
			"METADATA": CfnGlobalResolverGlobalResolverLogsRecordFields_METADATA,
			"SEVERITY_ID": CfnGlobalResolverGlobalResolverLogsRecordFields_SEVERITY_ID,
			"SRC_ENDPOINT": CfnGlobalResolverGlobalResolverLogsRecordFields_SRC_ENDPOINT,
			"TIME": CfnGlobalResolverGlobalResolverLogsRecordFields_TIME,
			"TYPE_UID": CfnGlobalResolverGlobalResolverLogsRecordFields_TYPE_UID,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsS3Props",
		reflect.TypeOf((*CfnGlobalResolverGlobalResolverLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverLogsMixin",
		reflect.TypeOf((*CfnGlobalResolverLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGlobalResolverLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
