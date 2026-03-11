package previewawstransfermixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerLogsMixin",
		reflect.TypeOf((*CfnServerLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServerLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerTransferLogs",
		reflect.TypeOf((*CfnServerTransferLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnServerTransferLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerTransferLogsFirehoseProps",
		reflect.TypeOf((*CfnServerTransferLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerTransferLogsLogGroupProps",
		reflect.TypeOf((*CfnServerTransferLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerTransferLogsOutputFormat",
		reflect.TypeOf((*CfnServerTransferLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnServerTransferLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerTransferLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnServerTransferLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnServerTransferLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnServerTransferLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnServerTransferLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerTransferLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnServerTransferLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnServerTransferLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnServerTransferLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerTransferLogsOutputFormat.S3",
		reflect.TypeOf((*CfnServerTransferLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnServerTransferLogsOutputFormat_S3_JSON,
			"PLAIN": CfnServerTransferLogsOutputFormat_S3_PLAIN,
			"W3C": CfnServerTransferLogsOutputFormat_S3_W3C,
			"PARQUET": CfnServerTransferLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerTransferLogsS3Props",
		reflect.TypeOf((*CfnServerTransferLogsS3Props)(nil)).Elem(),
	)
}
