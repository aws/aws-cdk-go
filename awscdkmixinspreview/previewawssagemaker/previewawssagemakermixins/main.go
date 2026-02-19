package previewawssagemakermixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigMixinProps",
		reflect.TypeOf((*CfnAppImageConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppImageConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin.CodeEditorAppImageConfigProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_CodeEditorAppImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin.ContainerConfigProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_ContainerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin.CustomImageContainerEnvironmentVariableProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_CustomImageContainerEnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin.FileSystemConfigProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin.JupyterLabAppImageConfigProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_JupyterLabAppImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin.KernelGatewayImageConfigProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_KernelGatewayImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin.KernelSpecProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_KernelSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppMixinProps",
		reflect.TypeOf((*CfnAppMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppPropsMixin",
		reflect.TypeOf((*CfnAppPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppPropsMixin.ResourceSpecProperty",
		reflect.TypeOf((*CfnAppPropsMixin_ResourceSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin",
		reflect.TypeOf((*CfnClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.AlarmDetailsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AlarmDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.CapacitySizeConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_CapacitySizeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterAutoScalingConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterAutoScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterCapacityRequirementsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterCapacityRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterEbsVolumeConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterEbsVolumeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterFsxLustreConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterFsxLustreConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterFsxOpenZfsConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterFsxOpenZfsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterInstanceGroupProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterInstanceGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterInstanceStorageConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterInstanceStorageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterKubernetesConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterKubernetesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterKubernetesTaintProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterKubernetesTaintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterLifeCycleConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterLifeCycleConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterOrchestratorEksConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterOrchestratorEksConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterOrchestratorSlurmConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterOrchestratorSlurmConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterRestrictedInstanceGroupProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterRestrictedInstanceGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ClusterSlurmConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterSlurmConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.DeploymentConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_DeploymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.EnvironmentConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EnvironmentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.FSxLustreConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_FSxLustreConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.OrchestratorProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OrchestratorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.RollingUpdatePolicyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_RollingUpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.ScheduledUpdateConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScheduledUpdateConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.TieredStorageConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_TieredStorageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnCodeRepositoryMixinProps",
		reflect.TypeOf((*CfnCodeRepositoryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnCodeRepositoryPropsMixin",
		reflect.TypeOf((*CfnCodeRepositoryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeRepositoryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnCodeRepositoryPropsMixin.GitConfigProperty",
		reflect.TypeOf((*CfnCodeRepositoryPropsMixin_GitConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionMixinProps",
		reflect.TypeOf((*CfnDataQualityJobDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataQualityJobDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.BatchTransformInputProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_BatchTransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.ConstraintsResourceProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_ConstraintsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.CsvProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.DataQualityAppSpecificationProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_DataQualityAppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.DataQualityBaselineConfigProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_DataQualityBaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.DataQualityJobInputProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_DataQualityJobInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.DatasetFormatProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_DatasetFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.EndpointInputProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_EndpointInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.JsonProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_JsonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.MonitoringOutputConfigProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_MonitoringOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.MonitoringOutputProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_MonitoringOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.MonitoringResourcesProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_MonitoringResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.StatisticsResourceProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_StatisticsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDataQualityJobDefinitionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDeviceFleetMixinProps",
		reflect.TypeOf((*CfnDeviceFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDeviceFleetPropsMixin",
		reflect.TypeOf((*CfnDeviceFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeviceFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDeviceFleetPropsMixin.EdgeOutputConfigProperty",
		reflect.TypeOf((*CfnDeviceFleetPropsMixin_EdgeOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDeviceMixinProps",
		reflect.TypeOf((*CfnDeviceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDevicePropsMixin",
		reflect.TypeOf((*CfnDevicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDevicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDevicePropsMixin.DeviceProperty",
		reflect.TypeOf((*CfnDevicePropsMixin_DeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.AppLifecycleManagementProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_AppLifecycleManagementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.CodeEditorAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CodeEditorAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.CodeRepositoryProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.CustomFileSystemConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CustomFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.CustomImageProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CustomImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.CustomPosixUserConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CustomPosixUserConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.DefaultEbsStorageSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DefaultEbsStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.DefaultSpaceSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DefaultSpaceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.DefaultSpaceStorageSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DefaultSpaceStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.DockerSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DockerSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.DomainSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DomainSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.EFSFileSystemConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_EFSFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.FSxLustreFileSystemConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_FSxLustreFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.HiddenSageMakerImageProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_HiddenSageMakerImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.IdleSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_IdleSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.JupyterLabAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_JupyterLabAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.JupyterServerAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_JupyterServerAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.KernelGatewayAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_KernelGatewayAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.RSessionAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_RSessionAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.RStudioServerProAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_RStudioServerProAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.RStudioServerProDomainSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_RStudioServerProDomainSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.ResourceSpecProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ResourceSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.S3FileSystemConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_S3FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.SharingSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SharingSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.StudioWebPortalSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_StudioWebPortalSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.UnifiedStudioSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_UnifiedStudioSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin.UserSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_UserSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigMixinProps",
		reflect.TypeOf((*CfnEndpointConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEndpointConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.AsyncInferenceClientConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_AsyncInferenceClientConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.AsyncInferenceConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_AsyncInferenceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.AsyncInferenceNotificationConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_AsyncInferenceNotificationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.AsyncInferenceOutputConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_AsyncInferenceOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.CapacityReservationConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_CapacityReservationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.CaptureContentTypeHeaderProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_CaptureContentTypeHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.CaptureOptionProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_CaptureOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.ClarifyExplainerConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ClarifyExplainerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.ClarifyInferenceConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ClarifyInferenceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.ClarifyShapBaselineConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ClarifyShapBaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.ClarifyShapConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ClarifyShapConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.ClarifyTextConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ClarifyTextConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.DataCaptureConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_DataCaptureConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.ExplainerConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ExplainerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.ManagedInstanceScalingProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ManagedInstanceScalingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.ProductionVariantProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ProductionVariantProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.RoutingConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_RoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.ServerlessConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ServerlessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointMixinProps",
		reflect.TypeOf((*CfnEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointPropsMixin",
		reflect.TypeOf((*CfnEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointPropsMixin.AlarmProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_AlarmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointPropsMixin.AutoRollbackConfigProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_AutoRollbackConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointPropsMixin.BlueGreenUpdatePolicyProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_BlueGreenUpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointPropsMixin.CapacitySizeProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_CapacitySizeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointPropsMixin.DeploymentConfigProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_DeploymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointPropsMixin.RollingUpdatePolicyProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_RollingUpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointPropsMixin.TrafficRoutingConfigProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_TrafficRoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointPropsMixin.VariantPropertyProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_VariantPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupMixinProps",
		reflect.TypeOf((*CfnFeatureGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFeatureGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin.DataCatalogConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_DataCatalogConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin.FeatureDefinitionProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_FeatureDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin.OfflineStoreConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_OfflineStoreConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin.OnlineStoreConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_OnlineStoreConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin.OnlineStoreSecurityConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_OnlineStoreSecurityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin.S3StorageConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_S3StorageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin.ThroughputConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_ThroughputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnFeatureGroupPropsMixin.TtlDurationProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_TtlDurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnImageMixinProps",
		reflect.TypeOf((*CfnImageMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnImagePropsMixin",
		reflect.TypeOf((*CfnImagePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnImagePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnImageVersionMixinProps",
		reflect.TypeOf((*CfnImageVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnImageVersionPropsMixin",
		reflect.TypeOf((*CfnImageVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnImageVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentMixinProps",
		reflect.TypeOf((*CfnInferenceComponentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInferenceComponentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.AlarmProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_AlarmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.AutoRollbackConfigurationProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_AutoRollbackConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.DeployedImageProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_DeployedImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.InferenceComponentCapacitySizeProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentCapacitySizeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.InferenceComponentComputeResourceRequirementsProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentComputeResourceRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.InferenceComponentContainerSpecificationProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentContainerSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.InferenceComponentDeploymentConfigProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentDeploymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.InferenceComponentRollingUpdatePolicyProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentRollingUpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.InferenceComponentRuntimeConfigProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentRuntimeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.InferenceComponentSpecificationProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin.InferenceComponentStartupParametersProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentStartupParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentMixinProps",
		reflect.TypeOf((*CfnInferenceExperimentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInferenceExperimentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin.CaptureContentTypeHeaderProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_CaptureContentTypeHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin.DataStorageConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_DataStorageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin.EndpointMetadataProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_EndpointMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin.InferenceExperimentScheduleProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_InferenceExperimentScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin.ModelInfrastructureConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_ModelInfrastructureConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin.ModelVariantConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_ModelVariantConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin.RealTimeInferenceConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_RealTimeInferenceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin.ShadowModeConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_ShadowModeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin.ShadowModelVariantConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_ShadowModelVariantConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMlflowTrackingServerMixinProps",
		reflect.TypeOf((*CfnMlflowTrackingServerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMlflowTrackingServerPropsMixin",
		reflect.TypeOf((*CfnMlflowTrackingServerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMlflowTrackingServerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionMixinProps",
		reflect.TypeOf((*CfnModelBiasJobDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelBiasJobDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.BatchTransformInputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_BatchTransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.ConstraintsResourceProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_ConstraintsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.CsvProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.DatasetFormatProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_DatasetFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.EndpointInputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_EndpointInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.JsonProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_JsonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.ModelBiasAppSpecificationProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_ModelBiasAppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.ModelBiasBaselineConfigProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_ModelBiasBaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.ModelBiasJobInputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_ModelBiasJobInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.MonitoringGroundTruthS3InputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_MonitoringGroundTruthS3InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.MonitoringOutputConfigProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_MonitoringOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.MonitoringOutputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_MonitoringOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.MonitoringResourcesProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_MonitoringResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelBiasJobDefinitionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardMixinProps",
		reflect.TypeOf((*CfnModelCardMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin",
		reflect.TypeOf((*CfnModelCardPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelCardPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.AdditionalInformationProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_AdditionalInformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.BusinessDetailsProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_BusinessDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.ContainerProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ContainerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.ContentProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.EvaluationDetailProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_EvaluationDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.FunctionProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_FunctionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.InferenceEnvironmentProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_InferenceEnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.InferenceSpecificationProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_InferenceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.IntendedUsesProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_IntendedUsesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.MetricDataItemsProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_MetricDataItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.MetricGroupProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_MetricGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.ModelOverviewProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ModelOverviewProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.ModelPackageCreatorProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ModelPackageCreatorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.ModelPackageDetailsProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ModelPackageDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.ObjectiveFunctionProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ObjectiveFunctionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.SecurityConfigProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_SecurityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.SourceAlgorithmProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_SourceAlgorithmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.TrainingDetailsProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_TrainingDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.TrainingEnvironmentProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_TrainingEnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.TrainingHyperParameterProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_TrainingHyperParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.TrainingJobDetailsProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_TrainingJobDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.TrainingMetricProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_TrainingMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelCardPropsMixin.UserContextProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_UserContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionMixinProps",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelExplainabilityJobDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.BatchTransformInputProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_BatchTransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.ConstraintsResourceProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_ConstraintsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.CsvProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.DatasetFormatProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_DatasetFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.EndpointInputProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_EndpointInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.JsonProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_JsonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.ModelExplainabilityAppSpecificationProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_ModelExplainabilityAppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.ModelExplainabilityBaselineConfigProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_ModelExplainabilityBaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.ModelExplainabilityJobInputProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_ModelExplainabilityJobInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.MonitoringOutputConfigProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_MonitoringOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.MonitoringOutputProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_MonitoringOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.MonitoringResourcesProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_MonitoringResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelExplainabilityJobDefinitionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelMixinProps",
		reflect.TypeOf((*CfnModelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackageGroupMixinProps",
		reflect.TypeOf((*CfnModelPackageGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackageGroupPropsMixin",
		reflect.TypeOf((*CfnModelPackageGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelPackageGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackageMixinProps",
		reflect.TypeOf((*CfnModelPackageMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin",
		reflect.TypeOf((*CfnModelPackagePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelPackagePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.AdditionalInferenceSpecificationDefinitionProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_AdditionalInferenceSpecificationDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.BiasProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_BiasProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.DataSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.DriftCheckBaselinesProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DriftCheckBaselinesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.DriftCheckBiasProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DriftCheckBiasProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.DriftCheckExplainabilityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DriftCheckExplainabilityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.DriftCheckModelDataQualityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DriftCheckModelDataQualityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.DriftCheckModelQualityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DriftCheckModelQualityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ExplainabilityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ExplainabilityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.FileSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_FileSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.InferenceSpecificationProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_InferenceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.MetadataPropertiesProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_MetadataPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.MetricsSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_MetricsSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ModelAccessConfigProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelAccessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ModelCardProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelCardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ModelDataQualityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelDataQualityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ModelDataSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ModelMetricsProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelMetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ModelPackageContainerDefinitionProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelPackageContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ModelPackageStatusDetailsProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelPackageStatusDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ModelPackageStatusItemProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelPackageStatusItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ModelQualityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelQualityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.S3DataSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_S3DataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.S3ModelDataSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_S3ModelDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.SecurityConfigProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_SecurityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.SourceAlgorithmProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_SourceAlgorithmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.SourceAlgorithmSpecificationProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_SourceAlgorithmSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.TransformInputProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_TransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.TransformJobDefinitionProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_TransformJobDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.TransformOutputProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_TransformOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.TransformResourcesProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_TransformResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ValidationProfileProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ValidationProfileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin.ValidationSpecificationProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ValidationSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin",
		reflect.TypeOf((*CfnModelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin.ContainerDefinitionProperty",
		reflect.TypeOf((*CfnModelPropsMixin_ContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin.HubAccessConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_HubAccessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin.ImageConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_ImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin.InferenceExecutionConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_InferenceExecutionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin.ModelAccessConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_ModelAccessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin.ModelDataSourceProperty",
		reflect.TypeOf((*CfnModelPropsMixin_ModelDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin.MultiModelConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_MultiModelConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin.RepositoryAuthConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_RepositoryAuthConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin.S3DataSourceProperty",
		reflect.TypeOf((*CfnModelPropsMixin_S3DataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionMixinProps",
		reflect.TypeOf((*CfnModelQualityJobDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelQualityJobDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.BatchTransformInputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_BatchTransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.ConstraintsResourceProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_ConstraintsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.CsvProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.DatasetFormatProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_DatasetFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.EndpointInputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_EndpointInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.JsonProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_JsonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.ModelQualityAppSpecificationProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_ModelQualityAppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.ModelQualityBaselineConfigProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_ModelQualityBaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.ModelQualityJobInputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_ModelQualityJobInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.MonitoringGroundTruthS3InputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_MonitoringGroundTruthS3InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.MonitoringOutputConfigProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_MonitoringOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.MonitoringOutputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_MonitoringOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.MonitoringResourcesProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_MonitoringResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringScheduleMixinProps",
		reflect.TypeOf((*CfnMonitoringScheduleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMonitoringSchedulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.BaselineConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_BaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.BatchTransformInputProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_BatchTransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.ConstraintsResourceProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_ConstraintsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.CsvProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.DatasetFormatProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_DatasetFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.EndpointInputProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_EndpointInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.JsonProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_JsonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.MonitoringAppSpecificationProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringAppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.MonitoringExecutionSummaryProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringExecutionSummaryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.MonitoringInputProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.MonitoringJobDefinitionProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringJobDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.MonitoringOutputConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.MonitoringOutputProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.MonitoringResourcesProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.MonitoringScheduleConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringScheduleConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.ScheduleConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_ScheduleConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.StatisticsResourceProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_StatisticsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnNotebookInstanceLifecycleConfigMixinProps",
		reflect.TypeOf((*CfnNotebookInstanceLifecycleConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnNotebookInstanceLifecycleConfigPropsMixin",
		reflect.TypeOf((*CfnNotebookInstanceLifecycleConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNotebookInstanceLifecycleConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnNotebookInstanceLifecycleConfigPropsMixin.NotebookInstanceLifecycleHookProperty",
		reflect.TypeOf((*CfnNotebookInstanceLifecycleConfigPropsMixin_NotebookInstanceLifecycleHookProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnNotebookInstanceMixinProps",
		reflect.TypeOf((*CfnNotebookInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnNotebookInstancePropsMixin",
		reflect.TypeOf((*CfnNotebookInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNotebookInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnNotebookInstancePropsMixin.InstanceMetadataServiceConfigurationProperty",
		reflect.TypeOf((*CfnNotebookInstancePropsMixin_InstanceMetadataServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnPartnerAppMixinProps",
		reflect.TypeOf((*CfnPartnerAppMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnPartnerAppPropsMixin",
		reflect.TypeOf((*CfnPartnerAppPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPartnerAppPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnPartnerAppPropsMixin.PartnerAppConfigProperty",
		reflect.TypeOf((*CfnPartnerAppPropsMixin_PartnerAppConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnPartnerAppPropsMixin.PartnerAppMaintenanceConfigProperty",
		reflect.TypeOf((*CfnPartnerAppPropsMixin_PartnerAppMaintenanceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnPipelineMixinProps",
		reflect.TypeOf((*CfnPipelineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnPipelinePropsMixin",
		reflect.TypeOf((*CfnPipelinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipelinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnPipelinePropsMixin.ParallelismConfigurationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ParallelismConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnPipelinePropsMixin.PipelineDefinitionProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_PipelineDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnPipelinePropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobMixinProps",
		reflect.TypeOf((*CfnProcessingJobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin",
		reflect.TypeOf((*CfnProcessingJobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProcessingJobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.AppSpecificationProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_AppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.AthenaDatasetDefinitionProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_AthenaDatasetDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.DatasetDefinitionProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_DatasetDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.ExperimentConfigProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ExperimentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.FeatureStoreOutputProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_FeatureStoreOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.ProcessingInputsObjectProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ProcessingInputsObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.ProcessingOutputConfigProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ProcessingOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.ProcessingOutputsObjectProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ProcessingOutputsObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.ProcessingResourcesProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ProcessingResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.RedshiftDatasetDefinitionProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_RedshiftDatasetDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.S3InputProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_S3InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProjectPropsMixin.CfnStackParameterProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_CfnStackParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProjectPropsMixin.CfnTemplateProviderDetailProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_CfnTemplateProviderDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProjectPropsMixin.ProvisioningParameterProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProvisioningParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProjectPropsMixin.ServiceCatalogProvisionedProductDetailsProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ServiceCatalogProvisionedProductDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProjectPropsMixin.ServiceCatalogProvisioningDetailsProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ServiceCatalogProvisioningDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProjectPropsMixin.TemplateProviderDetailProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_TemplateProviderDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpaceMixinProps",
		reflect.TypeOf((*CfnSpaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin",
		reflect.TypeOf((*CfnSpacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSpacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.CodeRepositoryProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_CodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.CustomFileSystemProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_CustomFileSystemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.CustomImageProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_CustomImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.EFSFileSystemProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_EFSFileSystemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.EbsStorageSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_EbsStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.FSxLustreFileSystemProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_FSxLustreFileSystemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.JupyterServerAppSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_JupyterServerAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.KernelGatewayAppSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_KernelGatewayAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.OwnershipSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_OwnershipSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.ResourceSpecProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_ResourceSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.S3FileSystemProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_S3FileSystemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.SpaceAppLifecycleManagementProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceAppLifecycleManagementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.SpaceCodeEditorAppSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceCodeEditorAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.SpaceIdleSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceIdleSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.SpaceJupyterLabAppSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceJupyterLabAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.SpaceSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.SpaceSharingSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceSharingSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin.SpaceStorageSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnStudioLifecycleConfigMixinProps",
		reflect.TypeOf((*CfnStudioLifecycleConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnStudioLifecycleConfigPropsMixin",
		reflect.TypeOf((*CfnStudioLifecycleConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStudioLifecycleConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfileMixinProps",
		reflect.TypeOf((*CfnUserProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin",
		reflect.TypeOf((*CfnUserProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.AppLifecycleManagementProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_AppLifecycleManagementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.CodeEditorAppSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_CodeEditorAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.CodeRepositoryProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_CodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.CustomFileSystemConfigProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_CustomFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.CustomImageProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_CustomImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.CustomPosixUserConfigProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_CustomPosixUserConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.DefaultEbsStorageSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_DefaultEbsStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.DefaultSpaceStorageSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_DefaultSpaceStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.EFSFileSystemConfigProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_EFSFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.FSxLustreFileSystemConfigProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_FSxLustreFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.HiddenSageMakerImageProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_HiddenSageMakerImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.IdleSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_IdleSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.JupyterLabAppSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_JupyterLabAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.JupyterServerAppSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_JupyterServerAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.KernelGatewayAppSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_KernelGatewayAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.RStudioServerProAppSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_RStudioServerProAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.ResourceSpecProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_ResourceSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.S3FileSystemConfigProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_S3FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.SharingSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_SharingSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.StudioWebPortalSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_StudioWebPortalSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin.UserSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_UserSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogs",
		reflect.TypeOf((*CfnWorkteamActivityLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnWorkteamActivityLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamLogsMixin",
		reflect.TypeOf((*CfnWorkteamLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkteamLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamMixinProps",
		reflect.TypeOf((*CfnWorkteamMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamPropsMixin",
		reflect.TypeOf((*CfnWorkteamPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkteamPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamPropsMixin.CognitoMemberDefinitionProperty",
		reflect.TypeOf((*CfnWorkteamPropsMixin_CognitoMemberDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamPropsMixin.MemberDefinitionProperty",
		reflect.TypeOf((*CfnWorkteamPropsMixin_MemberDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamPropsMixin.NotificationConfigurationProperty",
		reflect.TypeOf((*CfnWorkteamPropsMixin_NotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamPropsMixin.OidcMemberDefinitionProperty",
		reflect.TypeOf((*CfnWorkteamPropsMixin_OidcMemberDefinitionProperty)(nil)).Elem(),
	)
}
