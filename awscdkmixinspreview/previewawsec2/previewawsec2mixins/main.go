package previewawsec2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogs",
		reflect.TypeOf((*CfnRouteServerPeerEventLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRouteServerPeerEventLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsDestProps",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsFirehoseProps",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsLogGroupProps",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsOutputFormat",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnRouteServerPeerEventLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnRouteServerPeerEventLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnRouteServerPeerEventLogsOutputFormat_Firehose_JSON,
			"RAW": CfnRouteServerPeerEventLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnRouteServerPeerEventLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnRouteServerPeerEventLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsOutputFormat.S3",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnRouteServerPeerEventLogsOutputFormat_S3_PLAIN,
			"JSON": CfnRouteServerPeerEventLogsOutputFormat_S3_JSON,
			"W3C": CfnRouteServerPeerEventLogsOutputFormat_S3_W3C,
			"PARQUET": CfnRouteServerPeerEventLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsRecordFields",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnRouteServerPeerEventLogsRecordFields_TIMESTAMP,
			"RESOURCE_ID": CfnRouteServerPeerEventLogsRecordFields_RESOURCE_ID,
			"STATUS": CfnRouteServerPeerEventLogsRecordFields_STATUS,
			"MESSAGE": CfnRouteServerPeerEventLogsRecordFields_MESSAGE,
			"RESOURCE_ARN": CfnRouteServerPeerEventLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnRouteServerPeerEventLogsRecordFields_EVENT_TIMESTAMP,
			"TYPE": CfnRouteServerPeerEventLogsRecordFields_TYPE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsS3Props",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerLogsMixin",
		reflect.TypeOf((*CfnRouteServerPeerLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRouteServerPeerLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
