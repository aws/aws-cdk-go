package previewawsecsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogs",
		reflect.TypeOf((*CfnClusterActionLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterActionLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsDestProps",
		reflect.TypeOf((*CfnClusterActionLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterActionLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterActionLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsOutputFormat",
		reflect.TypeOf((*CfnClusterActionLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterActionLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterActionLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterActionLogsOutputFormat_Firehose_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterActionLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterActionLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterActionLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterActionLogsOutputFormat_S3_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsRecordFields",
		reflect.TypeOf((*CfnClusterActionLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnClusterActionLogsRecordFields_TIMESTAMP,
			"RESOURCEARN": CfnClusterActionLogsRecordFields_RESOURCEARN,
			"ACTIONSOURCEID": CfnClusterActionLogsRecordFields_ACTIONSOURCEID,
			"LOGLEVEL": CfnClusterActionLogsRecordFields_LOGLEVEL,
			"EVENTTIMESTAMP": CfnClusterActionLogsRecordFields_EVENTTIMESTAMP,
			"DETAIL": CfnClusterActionLogsRecordFields_DETAIL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsS3Props",
		reflect.TypeOf((*CfnClusterActionLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterLogsMixin",
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
}
