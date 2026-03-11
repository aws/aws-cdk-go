package previewawskafkaconnectmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogs",
		reflect.TypeOf((*CfnConnectorApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnConnectorApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnConnectorApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnConnectorApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnConnectorApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnConnectorApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnConnectorApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnConnectorApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnConnectorApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnConnectorApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnConnectorApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnConnectorApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnConnectorApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnConnectorApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnConnectorApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnConnectorApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnConnectorApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnConnectorApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogsS3Props",
		reflect.TypeOf((*CfnConnectorApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorLogsMixin",
		reflect.TypeOf((*CfnConnectorLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectorLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
