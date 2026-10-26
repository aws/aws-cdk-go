package previewawsgameliftmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueLogsMixin",
		reflect.TypeOf((*CfnGameSessionQueueLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGameSessionQueueLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogs",
		reflect.TypeOf((*CfnGameSessionQueueQueuePlacementLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnGameSessionQueueQueuePlacementLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsDestProps",
		reflect.TypeOf((*CfnGameSessionQueueQueuePlacementLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsFirehoseProps",
		reflect.TypeOf((*CfnGameSessionQueueQueuePlacementLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsLogGroupProps",
		reflect.TypeOf((*CfnGameSessionQueueQueuePlacementLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsOutputFormat",
		reflect.TypeOf((*CfnGameSessionQueueQueuePlacementLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnGameSessionQueueQueuePlacementLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnGameSessionQueueQueuePlacementLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnGameSessionQueueQueuePlacementLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnGameSessionQueueQueuePlacementLogsOutputFormat_Firehose_JSON,
			"RAW": CfnGameSessionQueueQueuePlacementLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnGameSessionQueueQueuePlacementLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnGameSessionQueueQueuePlacementLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnGameSessionQueueQueuePlacementLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsOutputFormat.S3",
		reflect.TypeOf((*CfnGameSessionQueueQueuePlacementLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnGameSessionQueueQueuePlacementLogsOutputFormat_S3_PLAIN,
			"JSON": CfnGameSessionQueueQueuePlacementLogsOutputFormat_S3_JSON,
			"W3C": CfnGameSessionQueueQueuePlacementLogsOutputFormat_S3_W3C,
			"PARQUET": CfnGameSessionQueueQueuePlacementLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsRecordFields",
		reflect.TypeOf((*CfnGameSessionQueueQueuePlacementLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnGameSessionQueueQueuePlacementLogsRecordFields_TIMESTAMP,
			"PLACEMENT_ID": CfnGameSessionQueueQueuePlacementLogsRecordFields_PLACEMENT_ID,
			"EVENT_TYPE": CfnGameSessionQueueQueuePlacementLogsRecordFields_EVENT_TYPE,
			"SUMMARY": CfnGameSessionQueueQueuePlacementLogsRecordFields_SUMMARY,
			"DATA": CfnGameSessionQueueQueuePlacementLogsRecordFields_DATA,
			"RESOURCE_ARN": CfnGameSessionQueueQueuePlacementLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnGameSessionQueueQueuePlacementLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsS3Props",
		reflect.TypeOf((*CfnGameSessionQueueQueuePlacementLogsS3Props)(nil)).Elem(),
	)
}
