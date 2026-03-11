package previewawscleanroomsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogs",
		reflect.TypeOf((*CfnMembershipAnalysisLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMembershipAnalysisLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogsFirehoseProps",
		reflect.TypeOf((*CfnMembershipAnalysisLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogsLogGroupProps",
		reflect.TypeOf((*CfnMembershipAnalysisLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogsOutputFormat",
		reflect.TypeOf((*CfnMembershipAnalysisLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnMembershipAnalysisLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnMembershipAnalysisLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMembershipAnalysisLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnMembershipAnalysisLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnMembershipAnalysisLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnMembershipAnalysisLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnMembershipAnalysisLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnMembershipAnalysisLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogsOutputFormat.S3",
		reflect.TypeOf((*CfnMembershipAnalysisLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMembershipAnalysisLogsOutputFormat_S3_JSON,
			"PLAIN": CfnMembershipAnalysisLogsOutputFormat_S3_PLAIN,
			"W3C": CfnMembershipAnalysisLogsOutputFormat_S3_W3C,
			"PARQUET": CfnMembershipAnalysisLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogsS3Props",
		reflect.TypeOf((*CfnMembershipAnalysisLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipLogsMixin",
		reflect.TypeOf((*CfnMembershipLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMembershipLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
