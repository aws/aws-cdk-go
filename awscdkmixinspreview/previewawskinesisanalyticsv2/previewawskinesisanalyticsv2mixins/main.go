package previewawskinesisanalyticsv2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationCloudWatchLoggingOptionMixinProps",
		reflect.TypeOf((*CfnApplicationCloudWatchLoggingOptionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationCloudWatchLoggingOptionPropsMixin",
		reflect.TypeOf((*CfnApplicationCloudWatchLoggingOptionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationCloudWatchLoggingOptionPropsMixin.CloudWatchLoggingOptionProperty",
		reflect.TypeOf((*CfnApplicationCloudWatchLoggingOptionPropsMixin_CloudWatchLoggingOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationOutputMixinProps",
		reflect.TypeOf((*CfnApplicationOutputMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationOutputPropsMixin",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationOutputPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationOutputPropsMixin.DestinationSchemaProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_DestinationSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationOutputPropsMixin.KinesisFirehoseOutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_KinesisFirehoseOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationOutputPropsMixin.KinesisStreamsOutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_KinesisStreamsOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationOutputPropsMixin.LambdaOutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_LambdaOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationOutputPropsMixin.OutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.ApplicationCodeConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationCodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.ApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.ApplicationEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.ApplicationMaintenanceConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationMaintenanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.ApplicationRestoreConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationRestoreConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.ApplicationSnapshotConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationSnapshotConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.ApplicationSystemRollbackConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationSystemRollbackConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.CSVMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CSVMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.CatalogConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CatalogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.CheckpointConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CheckpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.CodeContentProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CodeContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.CustomArtifactConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CustomArtifactConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.DeployAsApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_DeployAsApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.EnvironmentPropertiesProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_EnvironmentPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.FlinkApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_FlinkApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.FlinkRunConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_FlinkRunConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.GlueDataCatalogConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_GlueDataCatalogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.InputLambdaProcessorProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputLambdaProcessorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.InputParallelismProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputParallelismProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.InputProcessingConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputProcessingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.InputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.InputSchemaProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.JSONMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_JSONMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.KinesisFirehoseInputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_KinesisFirehoseInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.KinesisStreamsInputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_KinesisStreamsInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.MappingParametersProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.MavenReferenceProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MavenReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.ParallelismConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ParallelismConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.PropertyGroupProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_PropertyGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.RecordColumnProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_RecordColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.RecordFormatProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_RecordFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.RunConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_RunConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.S3ContentBaseLocationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_S3ContentBaseLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.S3ContentLocationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_S3ContentLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.SqlApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_SqlApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.ZeppelinApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ZeppelinApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin.ZeppelinMonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ZeppelinMonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationReferenceDataSourceMixinProps",
		reflect.TypeOf((*CfnApplicationReferenceDataSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationReferenceDataSourcePropsMixin",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationReferenceDataSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationReferenceDataSourcePropsMixin.CSVMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_CSVMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationReferenceDataSourcePropsMixin.JSONMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_JSONMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationReferenceDataSourcePropsMixin.MappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_MappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationReferenceDataSourcePropsMixin.RecordColumnProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_RecordColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationReferenceDataSourcePropsMixin.RecordFormatProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_RecordFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationReferenceDataSourcePropsMixin.ReferenceDataSourceProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_ReferenceDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationReferenceDataSourcePropsMixin.ReferenceSchemaProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_ReferenceSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationReferenceDataSourcePropsMixin.S3ReferenceDataSourceProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_S3ReferenceDataSourceProperty)(nil)).Elem(),
	)
}
