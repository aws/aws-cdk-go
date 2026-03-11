package previewawsmskmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogs",
		reflect.TypeOf((*CfnClusterBrokerLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterBrokerLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterBrokerLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterBrokerLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogsOutputFormat",
		reflect.TypeOf((*CfnClusterBrokerLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterBrokerLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterBrokerLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterBrokerLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnClusterBrokerLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnClusterBrokerLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterBrokerLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterBrokerLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnClusterBrokerLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterBrokerLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterBrokerLogsOutputFormat_S3_JSON,
			"PLAIN": CfnClusterBrokerLogsOutputFormat_S3_PLAIN,
			"W3C": CfnClusterBrokerLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterBrokerLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogsS3Props",
		reflect.TypeOf((*CfnClusterBrokerLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterLogsMixin",
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
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogs",
		reflect.TypeOf((*CfnReplicatorApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnReplicatorApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsDestProps",
		reflect.TypeOf((*CfnReplicatorApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnReplicatorApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnReplicatorApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnReplicatorApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnReplicatorApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnReplicatorApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnReplicatorApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnReplicatorApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnReplicatorApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnReplicatorApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnReplicatorApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnReplicatorApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnReplicatorApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnReplicatorApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnReplicatorApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnReplicatorApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnReplicatorApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsRecordFields",
		reflect.TypeOf((*CfnReplicatorApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"REPLICATOR_NAME": CfnReplicatorApplicationLogsRecordFields_REPLICATOR_NAME,
			"REPLICATOR_ID": CfnReplicatorApplicationLogsRecordFields_REPLICATOR_ID,
			"EVENT_TIMESTAMP": CfnReplicatorApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"MESSAGE": CfnReplicatorApplicationLogsRecordFields_MESSAGE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsS3Props",
		reflect.TypeOf((*CfnReplicatorApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorLogsMixin",
		reflect.TypeOf((*CfnReplicatorLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReplicatorLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
