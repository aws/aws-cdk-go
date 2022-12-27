package awskinesisanalyticsv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication",
		reflect.TypeOf((*CfnApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationConfiguration", GoGetter: "ApplicationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "applicationDescription", GoGetter: "ApplicationDescription"},
			_jsii_.MemberProperty{JsiiProperty: "applicationMaintenanceConfiguration", GoGetter: "ApplicationMaintenanceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "applicationMode", GoGetter: "ApplicationMode"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "runConfiguration", GoGetter: "RunConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeEnvironment", GoGetter: "RuntimeEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "serviceExecutionRole", GoGetter: "ServiceExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.ApplicationCodeConfigurationProperty",
		reflect.TypeOf((*CfnApplication_ApplicationCodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.ApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplication_ApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.ApplicationMaintenanceConfigurationProperty",
		reflect.TypeOf((*CfnApplication_ApplicationMaintenanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.ApplicationRestoreConfigurationProperty",
		reflect.TypeOf((*CfnApplication_ApplicationRestoreConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.ApplicationSnapshotConfigurationProperty",
		reflect.TypeOf((*CfnApplication_ApplicationSnapshotConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.CSVMappingParametersProperty",
		reflect.TypeOf((*CfnApplication_CSVMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.CatalogConfigurationProperty",
		reflect.TypeOf((*CfnApplication_CatalogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.CheckpointConfigurationProperty",
		reflect.TypeOf((*CfnApplication_CheckpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.CodeContentProperty",
		reflect.TypeOf((*CfnApplication_CodeContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.CustomArtifactConfigurationProperty",
		reflect.TypeOf((*CfnApplication_CustomArtifactConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.DeployAsApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplication_DeployAsApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.EnvironmentPropertiesProperty",
		reflect.TypeOf((*CfnApplication_EnvironmentPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.FlinkApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplication_FlinkApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.FlinkRunConfigurationProperty",
		reflect.TypeOf((*CfnApplication_FlinkRunConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.GlueDataCatalogConfigurationProperty",
		reflect.TypeOf((*CfnApplication_GlueDataCatalogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.InputLambdaProcessorProperty",
		reflect.TypeOf((*CfnApplication_InputLambdaProcessorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.InputParallelismProperty",
		reflect.TypeOf((*CfnApplication_InputParallelismProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.InputProcessingConfigurationProperty",
		reflect.TypeOf((*CfnApplication_InputProcessingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.InputProperty",
		reflect.TypeOf((*CfnApplication_InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.InputSchemaProperty",
		reflect.TypeOf((*CfnApplication_InputSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.JSONMappingParametersProperty",
		reflect.TypeOf((*CfnApplication_JSONMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.KinesisFirehoseInputProperty",
		reflect.TypeOf((*CfnApplication_KinesisFirehoseInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.KinesisStreamsInputProperty",
		reflect.TypeOf((*CfnApplication_KinesisStreamsInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.MappingParametersProperty",
		reflect.TypeOf((*CfnApplication_MappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.MavenReferenceProperty",
		reflect.TypeOf((*CfnApplication_MavenReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplication_MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.ParallelismConfigurationProperty",
		reflect.TypeOf((*CfnApplication_ParallelismConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.PropertyGroupProperty",
		reflect.TypeOf((*CfnApplication_PropertyGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.RecordColumnProperty",
		reflect.TypeOf((*CfnApplication_RecordColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.RecordFormatProperty",
		reflect.TypeOf((*CfnApplication_RecordFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.RunConfigurationProperty",
		reflect.TypeOf((*CfnApplication_RunConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.S3ContentBaseLocationProperty",
		reflect.TypeOf((*CfnApplication_S3ContentBaseLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.S3ContentLocationProperty",
		reflect.TypeOf((*CfnApplication_S3ContentLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.SqlApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplication_SqlApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.VpcConfigurationProperty",
		reflect.TypeOf((*CfnApplication_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.ZeppelinApplicationConfigurationProperty",
		reflect.TypeOf((*CfnApplication_ZeppelinApplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplication.ZeppelinMonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplication_ZeppelinMonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationCloudWatchLoggingOption",
		reflect.TypeOf((*CfnApplicationCloudWatchLoggingOption)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cloudWatchLoggingOption", GoGetter: "CloudWatchLoggingOption"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationCloudWatchLoggingOption{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationCloudWatchLoggingOption.CloudWatchLoggingOptionProperty",
		reflect.TypeOf((*CfnApplicationCloudWatchLoggingOption_CloudWatchLoggingOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationCloudWatchLoggingOptionProps",
		reflect.TypeOf((*CfnApplicationCloudWatchLoggingOptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationOutput",
		reflect.TypeOf((*CfnApplicationOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "output", GoGetter: "Output"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationOutput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationOutput.DestinationSchemaProperty",
		reflect.TypeOf((*CfnApplicationOutput_DestinationSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationOutput.KinesisFirehoseOutputProperty",
		reflect.TypeOf((*CfnApplicationOutput_KinesisFirehoseOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationOutput.KinesisStreamsOutputProperty",
		reflect.TypeOf((*CfnApplicationOutput_KinesisStreamsOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationOutput.LambdaOutputProperty",
		reflect.TypeOf((*CfnApplicationOutput_LambdaOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationOutput.OutputProperty",
		reflect.TypeOf((*CfnApplicationOutput_OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationOutputProps",
		reflect.TypeOf((*CfnApplicationOutputProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationProps",
		reflect.TypeOf((*CfnApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSource",
		reflect.TypeOf((*CfnApplicationReferenceDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "referenceDataSource", GoGetter: "ReferenceDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationReferenceDataSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSource.CSVMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSource_CSVMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSource.JSONMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSource_JSONMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSource.MappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSource_MappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSource.RecordColumnProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSource_RecordColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSource.RecordFormatProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSource_RecordFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSource.ReferenceDataSourceProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSource_ReferenceDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSource.ReferenceSchemaProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSource_ReferenceSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSource.S3ReferenceDataSourceProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSource_S3ReferenceDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisanalyticsv2.CfnApplicationReferenceDataSourceProps",
		reflect.TypeOf((*CfnApplicationReferenceDataSourceProps)(nil)).Elem(),
	)
}
