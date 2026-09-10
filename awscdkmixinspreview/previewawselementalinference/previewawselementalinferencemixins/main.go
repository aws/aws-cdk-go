package previewawselementalinferencemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogs",
		reflect.TypeOf((*CfnFeedApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFeedApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsDestProps",
		reflect.TypeOf((*CfnFeedApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnFeedApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnFeedApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnFeedApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnFeedApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnFeedApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFeedApplicationLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnFeedApplicationLogsOutputFormat_Firehose_JSON,
			"RAW": CfnFeedApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnFeedApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFeedApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnFeedApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnFeedApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFeedApplicationLogsOutputFormat_S3_PLAIN,
			"JSON": CfnFeedApplicationLogsOutputFormat_S3_JSON,
			"W3C": CfnFeedApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFeedApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsRecordFields",
		reflect.TypeOf((*CfnFeedApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnFeedApplicationLogsRecordFields_TIMESTAMP,
			"RESOURCE_ID": CfnFeedApplicationLogsRecordFields_RESOURCE_ID,
			"RESOURCE_TYPE": CfnFeedApplicationLogsRecordFields_RESOURCE_TYPE,
			"DETAIL": CfnFeedApplicationLogsRecordFields_DETAIL,
			"RESOURCE_ARN": CfnFeedApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnFeedApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"LOG_LEVEL": CfnFeedApplicationLogsRecordFields_LOG_LEVEL,
			"MESSAGE": CfnFeedApplicationLogsRecordFields_MESSAGE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsS3Props",
		reflect.TypeOf((*CfnFeedApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedLogsMixin",
		reflect.TypeOf((*CfnFeedLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFeedLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
