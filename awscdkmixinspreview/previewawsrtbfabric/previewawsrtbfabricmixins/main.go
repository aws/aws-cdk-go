package previewawsrtbfabricmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogs",
		reflect.TypeOf((*CfnLinkApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnLinkApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsDestProps",
		reflect.TypeOf((*CfnLinkApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnLinkApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnLinkApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnLinkApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnLinkApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnLinkApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnLinkApplicationLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnLinkApplicationLogsOutputFormat_Firehose_JSON,
			"RAW": CfnLinkApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnLinkApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnLinkApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnLinkApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnLinkApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnLinkApplicationLogsOutputFormat_S3_PLAIN,
			"JSON": CfnLinkApplicationLogsOutputFormat_S3_JSON,
			"W3C": CfnLinkApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnLinkApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsRecordFields",
		reflect.TypeOf((*CfnLinkApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnLinkApplicationLogsRecordFields_TIMESTAMP,
			"RESOURCE_ARN": CfnLinkApplicationLogsRecordFields_RESOURCE_ARN,
			"RESOURCE_ID": CfnLinkApplicationLogsRecordFields_RESOURCE_ID,
			"EVENT_TIMESTAMP": CfnLinkApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"BID_REQUEST_ID": CfnLinkApplicationLogsRecordFields_BID_REQUEST_ID,
			"RTBFABRIC_REQUEST_ID": CfnLinkApplicationLogsRecordFields_RTBFABRIC_REQUEST_ID,
			"MESSAGE": CfnLinkApplicationLogsRecordFields_MESSAGE,
			"STATUS_CODE": CfnLinkApplicationLogsRecordFields_STATUS_CODE,
			"CORRELATION_ID": CfnLinkApplicationLogsRecordFields_CORRELATION_ID,
			"RTBFABRIC_EVENT_TIMESTAMP": CfnLinkApplicationLogsRecordFields_RTBFABRIC_EVENT_TIMESTAMP,
			"EVENT_TYPE": CfnLinkApplicationLogsRecordFields_EVENT_TYPE,
			"LINK_ID": CfnLinkApplicationLogsRecordFields_LINK_ID,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsS3Props",
		reflect.TypeOf((*CfnLinkApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkLogsMixin",
		reflect.TypeOf((*CfnLinkLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLinkLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
