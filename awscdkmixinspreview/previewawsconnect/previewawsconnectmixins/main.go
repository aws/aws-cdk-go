package previewawsconnectmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogs",
		reflect.TypeOf((*CfnInstanceAmazonConnectFlowLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnInstanceAmazonConnectFlowLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsDestProps",
		reflect.TypeOf((*CfnInstanceAmazonConnectFlowLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsFirehoseProps",
		reflect.TypeOf((*CfnInstanceAmazonConnectFlowLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsLogGroupProps",
		reflect.TypeOf((*CfnInstanceAmazonConnectFlowLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsOutputFormat",
		reflect.TypeOf((*CfnInstanceAmazonConnectFlowLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnInstanceAmazonConnectFlowLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnInstanceAmazonConnectFlowLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnInstanceAmazonConnectFlowLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnInstanceAmazonConnectFlowLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnInstanceAmazonConnectFlowLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnInstanceAmazonConnectFlowLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnInstanceAmazonConnectFlowLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnInstanceAmazonConnectFlowLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsOutputFormat.S3",
		reflect.TypeOf((*CfnInstanceAmazonConnectFlowLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnInstanceAmazonConnectFlowLogsOutputFormat_S3_JSON,
			"PLAIN": CfnInstanceAmazonConnectFlowLogsOutputFormat_S3_PLAIN,
			"W3C": CfnInstanceAmazonConnectFlowLogsOutputFormat_S3_W3C,
			"PARQUET": CfnInstanceAmazonConnectFlowLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsRecordFields",
		reflect.TypeOf((*CfnInstanceAmazonConnectFlowLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"INSTANCEARN": CfnInstanceAmazonConnectFlowLogsRecordFields_INSTANCEARN,
			"TIMESTAMP": CfnInstanceAmazonConnectFlowLogsRecordFields_TIMESTAMP,
			"ACCOUNTID": CfnInstanceAmazonConnectFlowLogsRecordFields_ACCOUNTID,
			"RESOURCEARN": CfnInstanceAmazonConnectFlowLogsRecordFields_RESOURCEARN,
			"RESOURCETYPE": CfnInstanceAmazonConnectFlowLogsRecordFields_RESOURCETYPE,
			"RESOURCENAME": CfnInstanceAmazonConnectFlowLogsRecordFields_RESOURCENAME,
			"RESOURCEEXECUTIONSTARTTIME": CfnInstanceAmazonConnectFlowLogsRecordFields_RESOURCEEXECUTIONSTARTTIME,
			"RESOURCEPUBLISHTIMESTAMP": CfnInstanceAmazonConnectFlowLogsRecordFields_RESOURCEPUBLISHTIMESTAMP,
			"FLOWTYPE": CfnInstanceAmazonConnectFlowLogsRecordFields_FLOWTYPE,
			"CONTACTID": CfnInstanceAmazonConnectFlowLogsRecordFields_CONTACTID,
			"INITIATIONMETHOD": CfnInstanceAmazonConnectFlowLogsRecordFields_INITIATIONMETHOD,
			"CHANNEL": CfnInstanceAmazonConnectFlowLogsRecordFields_CHANNEL,
			"SUBTYPE": CfnInstanceAmazonConnectFlowLogsRecordFields_SUBTYPE,
			"LOGVERSION": CfnInstanceAmazonConnectFlowLogsRecordFields_LOGVERSION,
			"LOGTYPE": CfnInstanceAmazonConnectFlowLogsRecordFields_LOGTYPE,
			"ACTIONLOGDATA": CfnInstanceAmazonConnectFlowLogsRecordFields_ACTIONLOGDATA,
			"MODULECALLSTACK": CfnInstanceAmazonConnectFlowLogsRecordFields_MODULECALLSTACK,
			"ACTIONCOUNTER": CfnInstanceAmazonConnectFlowLogsRecordFields_ACTIONCOUNTER,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsS3Props",
		reflect.TypeOf((*CfnInstanceAmazonConnectFlowLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceLogsMixin",
		reflect.TypeOf((*CfnInstanceLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstanceLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
