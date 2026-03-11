package awskinesisanalyticsv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationCloudWatchLoggingOptionMixinProps",
		reflect.TypeOf((*CfnApplicationCloudWatchLoggingOptionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationCloudWatchLoggingOptionPropsMixin",
		reflect.TypeOf((*CfnApplicationCloudWatchLoggingOptionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationCloudWatchLoggingOptionPropsMixin.CloudWatchLoggingOptionProperty",
		reflect.TypeOf((*CfnApplicationCloudWatchLoggingOptionPropsMixin_CloudWatchLoggingOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationOutputMixinProps",
		reflect.TypeOf((*CfnApplicationOutputMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationOutputPropsMixin",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationOutputPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationOutputPropsMixin.DestinationSchemaProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_DestinationSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationOutputPropsMixin.KinesisFirehoseOutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_KinesisFirehoseOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationOutputPropsMixin.KinesisStreamsOutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_KinesisStreamsOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationOutputPropsMixin.LambdaOutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_LambdaOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationOutputPropsMixin.OutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.ApplicationCodeConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationCodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.ApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.ApplicationEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.ApplicationMaintenanceConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationMaintenanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.ApplicationRestoreConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationRestoreConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.ApplicationSnapshotConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationSnapshotConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.ApplicationSystemRollbackConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationSystemRollbackConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.CSVMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CSVMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.CatalogConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CatalogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.CheckpointConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CheckpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.CodeContentProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CodeContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.CustomArtifactConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CustomArtifactConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.DeployAsApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_DeployAsApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.EnvironmentPropertiesProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_EnvironmentPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.FlinkApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_FlinkApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.FlinkRunConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_FlinkRunConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.GlueDataCatalogConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_GlueDataCatalogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.InputLambdaProcessorProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputLambdaProcessorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.InputParallelismProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputParallelismProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.InputProcessingConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputProcessingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.InputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.InputSchemaProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.JSONMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_JSONMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.KinesisFirehoseInputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_KinesisFirehoseInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.KinesisStreamsInputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_KinesisStreamsInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.MappingParametersProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.MavenReferenceProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MavenReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.ParallelismConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ParallelismConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.PropertyGroupProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_PropertyGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.RecordColumnProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_RecordColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.RecordFormatProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_RecordFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.RunConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_RunConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.S3ContentBaseLocationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_S3ContentBaseLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.S3ContentLocationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_S3ContentLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.SqlApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_SqlApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.ZeppelinApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ZeppelinApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationPropsMixin.ZeppelinMonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ZeppelinMonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourceMixinProps",
		reflect.TypeOf((*CfnApplicationReferenceDataSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourcePropsMixin",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationReferenceDataSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourcePropsMixin.CSVMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_CSVMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourcePropsMixin.JSONMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_JSONMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourcePropsMixin.MappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_MappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourcePropsMixin.RecordColumnProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_RecordColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourcePropsMixin.RecordFormatProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_RecordFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourcePropsMixin.ReferenceDataSourceProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_ReferenceDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourcePropsMixin.ReferenceSchemaProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_ReferenceSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourcePropsMixin.S3ReferenceDataSourceProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_S3ReferenceDataSourceProperty)(nil)).Elem(),
	)
}
