package previewawsm2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationBatchJobLogs",
		reflect.TypeOf((*CfnApplicationBatchJobLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnApplicationBatchJobLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationBatchJobLogsFirehoseProps",
		reflect.TypeOf((*CfnApplicationBatchJobLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationBatchJobLogsLogGroupProps",
		reflect.TypeOf((*CfnApplicationBatchJobLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationBatchJobLogsOutputFormat",
		reflect.TypeOf((*CfnApplicationBatchJobLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnApplicationBatchJobLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationBatchJobLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnApplicationBatchJobLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationBatchJobLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnApplicationBatchJobLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnApplicationBatchJobLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationBatchJobLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnApplicationBatchJobLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnApplicationBatchJobLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnApplicationBatchJobLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationBatchJobLogsOutputFormat.S3",
		reflect.TypeOf((*CfnApplicationBatchJobLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationBatchJobLogsOutputFormat_S3_JSON,
			"PLAIN": CfnApplicationBatchJobLogsOutputFormat_S3_PLAIN,
			"W3C": CfnApplicationBatchJobLogsOutputFormat_S3_W3C,
			"PARQUET": CfnApplicationBatchJobLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationBatchJobLogsS3Props",
		reflect.TypeOf((*CfnApplicationBatchJobLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConfigLogs",
		reflect.TypeOf((*CfnApplicationConfigLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnApplicationConfigLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConfigLogsFirehoseProps",
		reflect.TypeOf((*CfnApplicationConfigLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConfigLogsLogGroupProps",
		reflect.TypeOf((*CfnApplicationConfigLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConfigLogsOutputFormat",
		reflect.TypeOf((*CfnApplicationConfigLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnApplicationConfigLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConfigLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnApplicationConfigLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationConfigLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnApplicationConfigLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnApplicationConfigLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConfigLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnApplicationConfigLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnApplicationConfigLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnApplicationConfigLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConfigLogsOutputFormat.S3",
		reflect.TypeOf((*CfnApplicationConfigLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationConfigLogsOutputFormat_S3_JSON,
			"PLAIN": CfnApplicationConfigLogsOutputFormat_S3_PLAIN,
			"W3C": CfnApplicationConfigLogsOutputFormat_S3_W3C,
			"PARQUET": CfnApplicationConfigLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConfigLogsS3Props",
		reflect.TypeOf((*CfnApplicationConfigLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConsoleLogs",
		reflect.TypeOf((*CfnApplicationConsoleLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnApplicationConsoleLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConsoleLogsFirehoseProps",
		reflect.TypeOf((*CfnApplicationConsoleLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConsoleLogsLogGroupProps",
		reflect.TypeOf((*CfnApplicationConsoleLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConsoleLogsOutputFormat",
		reflect.TypeOf((*CfnApplicationConsoleLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnApplicationConsoleLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConsoleLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnApplicationConsoleLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationConsoleLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnApplicationConsoleLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnApplicationConsoleLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConsoleLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnApplicationConsoleLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnApplicationConsoleLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnApplicationConsoleLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConsoleLogsOutputFormat.S3",
		reflect.TypeOf((*CfnApplicationConsoleLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationConsoleLogsOutputFormat_S3_JSON,
			"PLAIN": CfnApplicationConsoleLogsOutputFormat_S3_PLAIN,
			"W3C": CfnApplicationConsoleLogsOutputFormat_S3_W3C,
			"PARQUET": CfnApplicationConsoleLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConsoleLogsS3Props",
		reflect.TypeOf((*CfnApplicationConsoleLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogs",
		reflect.TypeOf((*CfnApplicationDatasetImportLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnApplicationDatasetImportLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogsFirehoseProps",
		reflect.TypeOf((*CfnApplicationDatasetImportLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogsLogGroupProps",
		reflect.TypeOf((*CfnApplicationDatasetImportLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogsOutputFormat",
		reflect.TypeOf((*CfnApplicationDatasetImportLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnApplicationDatasetImportLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnApplicationDatasetImportLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationDatasetImportLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnApplicationDatasetImportLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnApplicationDatasetImportLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnApplicationDatasetImportLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnApplicationDatasetImportLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnApplicationDatasetImportLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogsOutputFormat.S3",
		reflect.TypeOf((*CfnApplicationDatasetImportLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationDatasetImportLogsOutputFormat_S3_JSON,
			"PLAIN": CfnApplicationDatasetImportLogsOutputFormat_S3_PLAIN,
			"W3C": CfnApplicationDatasetImportLogsOutputFormat_S3_W3C,
			"PARQUET": CfnApplicationDatasetImportLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogsS3Props",
		reflect.TypeOf((*CfnApplicationDatasetImportLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationLogsMixin",
		reflect.TypeOf((*CfnApplicationLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationPropsMixin.DefinitionProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_DefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnDeploymentMixinProps",
		reflect.TypeOf((*CfnDeploymentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnDeploymentPropsMixin",
		reflect.TypeOf((*CfnDeploymentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnEnvironmentMixinProps",
		reflect.TypeOf((*CfnEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnEnvironmentPropsMixin",
		reflect.TypeOf((*CfnEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnEnvironmentPropsMixin.EfsStorageConfigurationProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_EfsStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnEnvironmentPropsMixin.FsxStorageConfigurationProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_FsxStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnEnvironmentPropsMixin.HighAvailabilityConfigProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_HighAvailabilityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnEnvironmentPropsMixin.StorageConfigurationProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_StorageConfigurationProperty)(nil)).Elem(),
	)
}
