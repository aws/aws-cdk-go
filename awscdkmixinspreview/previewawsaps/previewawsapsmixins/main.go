package previewawsapsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperApplicationLogs",
		reflect.TypeOf((*CfnScraperApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnScraperApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnScraperApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnScraperApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnScraperApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnScraperApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnScraperApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnScraperApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnScraperApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnScraperApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnScraperApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnScraperApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnScraperApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnScraperApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnScraperApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnScraperApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnScraperApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnScraperApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperApplicationLogsS3Props",
		reflect.TypeOf((*CfnScraperApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperLogsMixin",
		reflect.TypeOf((*CfnScraperLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScraperLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceLogsMixin",
		reflect.TypeOf((*CfnWorkspaceLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkspaceLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceManagedPrometheusLogs",
		reflect.TypeOf((*CfnWorkspaceManagedPrometheusLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnWorkspaceManagedPrometheusLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceManagedPrometheusLogsFirehoseProps",
		reflect.TypeOf((*CfnWorkspaceManagedPrometheusLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceManagedPrometheusLogsLogGroupProps",
		reflect.TypeOf((*CfnWorkspaceManagedPrometheusLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceManagedPrometheusLogsOutputFormat",
		reflect.TypeOf((*CfnWorkspaceManagedPrometheusLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnWorkspaceManagedPrometheusLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceManagedPrometheusLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnWorkspaceManagedPrometheusLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnWorkspaceManagedPrometheusLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnWorkspaceManagedPrometheusLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnWorkspaceManagedPrometheusLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceManagedPrometheusLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnWorkspaceManagedPrometheusLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnWorkspaceManagedPrometheusLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnWorkspaceManagedPrometheusLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceManagedPrometheusLogsOutputFormat.S3",
		reflect.TypeOf((*CfnWorkspaceManagedPrometheusLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnWorkspaceManagedPrometheusLogsOutputFormat_S3_JSON,
			"PLAIN": CfnWorkspaceManagedPrometheusLogsOutputFormat_S3_PLAIN,
			"W3C": CfnWorkspaceManagedPrometheusLogsOutputFormat_S3_W3C,
			"PARQUET": CfnWorkspaceManagedPrometheusLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceManagedPrometheusLogsS3Props",
		reflect.TypeOf((*CfnWorkspaceManagedPrometheusLogsS3Props)(nil)).Elem(),
	)
}
