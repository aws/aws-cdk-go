package awssagemaker

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppImageConfigMixinProps",
		reflect.TypeOf((*CfnAppImageConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppImageConfigPropsMixin",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppImageConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppImageConfigPropsMixin.CodeEditorAppImageConfigProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_CodeEditorAppImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppImageConfigPropsMixin.ContainerConfigProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_ContainerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppImageConfigPropsMixin.CustomImageContainerEnvironmentVariableProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_CustomImageContainerEnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppImageConfigPropsMixin.FileSystemConfigProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppImageConfigPropsMixin.JupyterLabAppImageConfigProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_JupyterLabAppImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppImageConfigPropsMixin.KernelGatewayImageConfigProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_KernelGatewayImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppImageConfigPropsMixin.KernelSpecProperty",
		reflect.TypeOf((*CfnAppImageConfigPropsMixin_KernelSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppMixinProps",
		reflect.TypeOf((*CfnAppMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppPropsMixin",
		reflect.TypeOf((*CfnAppPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAppPropsMixin.ResourceSpecProperty",
		reflect.TypeOf((*CfnAppPropsMixin_ResourceSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin",
		reflect.TypeOf((*CfnClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.AlarmDetailsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AlarmDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.CapacitySizeConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_CapacitySizeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterAutoScalingConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterAutoScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterCapacityRequirementsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterCapacityRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterEbsVolumeConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterEbsVolumeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterFsxLustreConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterFsxLustreConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterFsxOpenZfsConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterFsxOpenZfsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterInstanceGroupProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterInstanceGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterInstanceStorageConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterInstanceStorageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterKubernetesConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterKubernetesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterKubernetesTaintProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterKubernetesTaintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterLifeCycleConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterLifeCycleConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterNetworkInterfaceProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterNetworkInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterOrchestratorEksConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterOrchestratorEksConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterOrchestratorSlurmConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterOrchestratorSlurmConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterRestrictedInstanceGroupProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterRestrictedInstanceGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ClusterSlurmConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterSlurmConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.DeploymentConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_DeploymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.EnvironmentConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EnvironmentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.FSxLustreConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_FSxLustreConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.InstanceRequirementsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.OrchestratorProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OrchestratorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.RollingUpdatePolicyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_RollingUpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.ScheduledUpdateConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScheduledUpdateConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.TieredStorageConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_TieredStorageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnClusterPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnCodeRepositoryMixinProps",
		reflect.TypeOf((*CfnCodeRepositoryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnCodeRepositoryPropsMixin",
		reflect.TypeOf((*CfnCodeRepositoryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeRepositoryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnCodeRepositoryPropsMixin.GitConfigProperty",
		reflect.TypeOf((*CfnCodeRepositoryPropsMixin_GitConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionMixinProps",
		reflect.TypeOf((*CfnDataQualityJobDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataQualityJobDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.BatchTransformInputProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_BatchTransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.ConstraintsResourceProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_ConstraintsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.CsvProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.DataQualityAppSpecificationProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_DataQualityAppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.DataQualityBaselineConfigProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_DataQualityBaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.DataQualityJobInputProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_DataQualityJobInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.DatasetFormatProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_DatasetFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.EndpointInputProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_EndpointInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.JsonProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_JsonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.MonitoringOutputConfigProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_MonitoringOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.MonitoringOutputProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_MonitoringOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.MonitoringResourcesProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_MonitoringResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.StatisticsResourceProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_StatisticsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDataQualityJobDefinitionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnDataQualityJobDefinitionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDeviceFleetMixinProps",
		reflect.TypeOf((*CfnDeviceFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDeviceFleetPropsMixin",
		reflect.TypeOf((*CfnDeviceFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeviceFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDeviceFleetPropsMixin.EdgeOutputConfigProperty",
		reflect.TypeOf((*CfnDeviceFleetPropsMixin_EdgeOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDeviceMixinProps",
		reflect.TypeOf((*CfnDeviceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDevicePropsMixin",
		reflect.TypeOf((*CfnDevicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDevicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDevicePropsMixin.DeviceProperty",
		reflect.TypeOf((*CfnDevicePropsMixin_DeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.AppLifecycleManagementProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_AppLifecycleManagementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.CodeEditorAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CodeEditorAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.CodeRepositoryProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.CustomFileSystemConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CustomFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.CustomImageProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CustomImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.CustomPosixUserConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CustomPosixUserConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.DefaultEbsStorageSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DefaultEbsStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.DefaultSpaceSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DefaultSpaceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.DefaultSpaceStorageSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DefaultSpaceStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.DockerSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DockerSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.DomainSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DomainSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.EFSFileSystemConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_EFSFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.FSxLustreFileSystemConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_FSxLustreFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.HiddenSageMakerImageProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_HiddenSageMakerImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.IdleSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_IdleSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.JupyterLabAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_JupyterLabAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.JupyterServerAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_JupyterServerAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.KernelGatewayAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_KernelGatewayAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.RSessionAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_RSessionAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.RStudioServerProAppSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_RStudioServerProAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.RStudioServerProDomainSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_RStudioServerProDomainSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.ResourceSpecProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ResourceSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.S3FileSystemConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_S3FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.SharingSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SharingSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.StudioWebPortalSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_StudioWebPortalSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.UnifiedStudioSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_UnifiedStudioSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDomainPropsMixin.UserSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_UserSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigMixinProps",
		reflect.TypeOf((*CfnEndpointConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEndpointConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.AsyncInferenceClientConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_AsyncInferenceClientConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.AsyncInferenceConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_AsyncInferenceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.AsyncInferenceNotificationConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_AsyncInferenceNotificationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.AsyncInferenceOutputConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_AsyncInferenceOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.CapacityReservationConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_CapacityReservationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.CaptureContentTypeHeaderProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_CaptureContentTypeHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.CaptureOptionProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_CaptureOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.ClarifyExplainerConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ClarifyExplainerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.ClarifyInferenceConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ClarifyInferenceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.ClarifyShapBaselineConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ClarifyShapBaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.ClarifyShapConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ClarifyShapConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.ClarifyTextConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ClarifyTextConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.DataCaptureConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_DataCaptureConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.ExplainerConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ExplainerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.InstancePoolsProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_InstancePoolsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.ManagedInstanceScalingProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ManagedInstanceScalingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.ProductionVariantProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ProductionVariantProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.RoutingConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_RoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.ServerlessConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_ServerlessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointConfigPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnEndpointConfigPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointMixinProps",
		reflect.TypeOf((*CfnEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointPropsMixin",
		reflect.TypeOf((*CfnEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointPropsMixin.AlarmProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_AlarmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointPropsMixin.AutoRollbackConfigProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_AutoRollbackConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointPropsMixin.BlueGreenUpdatePolicyProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_BlueGreenUpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointPropsMixin.CapacitySizeProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_CapacitySizeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointPropsMixin.DeploymentConfigProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_DeploymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointPropsMixin.RollingUpdatePolicyProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_RollingUpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointPropsMixin.TrafficRoutingConfigProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_TrafficRoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnEndpointPropsMixin.VariantPropertyProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_VariantPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnFeatureGroupMixinProps",
		reflect.TypeOf((*CfnFeatureGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnFeatureGroupPropsMixin",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFeatureGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnFeatureGroupPropsMixin.DataCatalogConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_DataCatalogConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnFeatureGroupPropsMixin.FeatureDefinitionProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_FeatureDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnFeatureGroupPropsMixin.OfflineStoreConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_OfflineStoreConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnFeatureGroupPropsMixin.OnlineStoreConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_OnlineStoreConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnFeatureGroupPropsMixin.OnlineStoreSecurityConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_OnlineStoreSecurityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnFeatureGroupPropsMixin.S3StorageConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_S3StorageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnFeatureGroupPropsMixin.ThroughputConfigProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_ThroughputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnFeatureGroupPropsMixin.TtlDurationProperty",
		reflect.TypeOf((*CfnFeatureGroupPropsMixin_TtlDurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnImageMixinProps",
		reflect.TypeOf((*CfnImageMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnImagePropsMixin",
		reflect.TypeOf((*CfnImagePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnImagePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnImageVersionMixinProps",
		reflect.TypeOf((*CfnImageVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnImageVersionPropsMixin",
		reflect.TypeOf((*CfnImageVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnImageVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentMixinProps",
		reflect.TypeOf((*CfnInferenceComponentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInferenceComponentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.AlarmProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_AlarmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.AutoRollbackConfigurationProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_AutoRollbackConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.DeployedImageProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_DeployedImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.InferenceComponentCapacitySizeProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentCapacitySizeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.InferenceComponentComputeResourceRequirementsProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentComputeResourceRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.InferenceComponentContainerSpecificationProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentContainerSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.InferenceComponentDeploymentConfigProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentDeploymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.InferenceComponentRollingUpdatePolicyProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentRollingUpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.InferenceComponentRuntimeConfigProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentRuntimeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.InferenceComponentSpecificationProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceComponentPropsMixin.InferenceComponentStartupParametersProperty",
		reflect.TypeOf((*CfnInferenceComponentPropsMixin_InferenceComponentStartupParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentMixinProps",
		reflect.TypeOf((*CfnInferenceExperimentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentPropsMixin",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInferenceExperimentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentPropsMixin.CaptureContentTypeHeaderProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_CaptureContentTypeHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentPropsMixin.DataStorageConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_DataStorageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentPropsMixin.EndpointMetadataProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_EndpointMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentPropsMixin.InferenceExperimentScheduleProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_InferenceExperimentScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentPropsMixin.ModelInfrastructureConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_ModelInfrastructureConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentPropsMixin.ModelVariantConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_ModelVariantConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentPropsMixin.RealTimeInferenceConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_RealTimeInferenceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentPropsMixin.ShadowModeConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_ShadowModeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnInferenceExperimentPropsMixin.ShadowModelVariantConfigProperty",
		reflect.TypeOf((*CfnInferenceExperimentPropsMixin_ShadowModelVariantConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMlflowAppMixinProps",
		reflect.TypeOf((*CfnMlflowAppMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMlflowAppPropsMixin",
		reflect.TypeOf((*CfnMlflowAppPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMlflowAppPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMlflowTrackingServerMixinProps",
		reflect.TypeOf((*CfnMlflowTrackingServerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMlflowTrackingServerPropsMixin",
		reflect.TypeOf((*CfnMlflowTrackingServerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMlflowTrackingServerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionMixinProps",
		reflect.TypeOf((*CfnModelBiasJobDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelBiasJobDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.BatchTransformInputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_BatchTransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.ConstraintsResourceProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_ConstraintsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.CsvProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.DatasetFormatProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_DatasetFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.EndpointInputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_EndpointInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.JsonProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_JsonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.ModelBiasAppSpecificationProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_ModelBiasAppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.ModelBiasBaselineConfigProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_ModelBiasBaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.ModelBiasJobInputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_ModelBiasJobInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.MonitoringGroundTruthS3InputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_MonitoringGroundTruthS3InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.MonitoringOutputConfigProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_MonitoringOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.MonitoringOutputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_MonitoringOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.MonitoringResourcesProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_MonitoringResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelBiasJobDefinitionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnModelBiasJobDefinitionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardMixinProps",
		reflect.TypeOf((*CfnModelCardMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin",
		reflect.TypeOf((*CfnModelCardPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelCardPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.AdditionalInformationProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_AdditionalInformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.BusinessDetailsProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_BusinessDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.ContainerProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ContainerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.ContentProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.EvaluationDetailProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_EvaluationDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.FunctionProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_FunctionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.InferenceEnvironmentProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_InferenceEnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.InferenceSpecificationProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_InferenceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.IntendedUsesProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_IntendedUsesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.MetricDataItemsProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_MetricDataItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.MetricGroupProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_MetricGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.ModelOverviewProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ModelOverviewProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.ModelPackageCreatorProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ModelPackageCreatorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.ModelPackageDetailsProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ModelPackageDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.ObjectiveFunctionProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_ObjectiveFunctionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.SecurityConfigProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_SecurityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.SourceAlgorithmProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_SourceAlgorithmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.TrainingDetailsProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_TrainingDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.TrainingEnvironmentProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_TrainingEnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.TrainingHyperParameterProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_TrainingHyperParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.TrainingJobDetailsProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_TrainingJobDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.TrainingMetricProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_TrainingMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelCardPropsMixin.UserContextProperty",
		reflect.TypeOf((*CfnModelCardPropsMixin_UserContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionMixinProps",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelExplainabilityJobDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.BatchTransformInputProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_BatchTransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.ConstraintsResourceProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_ConstraintsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.CsvProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.DatasetFormatProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_DatasetFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.EndpointInputProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_EndpointInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.JsonProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_JsonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.ModelExplainabilityAppSpecificationProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_ModelExplainabilityAppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.ModelExplainabilityBaselineConfigProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_ModelExplainabilityBaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.ModelExplainabilityJobInputProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_ModelExplainabilityJobInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.MonitoringOutputConfigProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_MonitoringOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.MonitoringOutputProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_MonitoringOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.MonitoringResourcesProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_MonitoringResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelExplainabilityJobDefinitionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnModelExplainabilityJobDefinitionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelMixinProps",
		reflect.TypeOf((*CfnModelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackageGroupMixinProps",
		reflect.TypeOf((*CfnModelPackageGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackageGroupPropsMixin",
		reflect.TypeOf((*CfnModelPackageGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelPackageGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackageMixinProps",
		reflect.TypeOf((*CfnModelPackageMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin",
		reflect.TypeOf((*CfnModelPackagePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelPackagePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.AdditionalInferenceSpecificationDefinitionProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_AdditionalInferenceSpecificationDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.BiasProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_BiasProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.DataSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.DriftCheckBaselinesProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DriftCheckBaselinesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.DriftCheckBiasProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DriftCheckBiasProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.DriftCheckExplainabilityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DriftCheckExplainabilityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.DriftCheckModelDataQualityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DriftCheckModelDataQualityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.DriftCheckModelQualityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_DriftCheckModelQualityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ExplainabilityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ExplainabilityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.FileSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_FileSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.InferenceSpecificationProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_InferenceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.MetadataPropertiesProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_MetadataPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.MetricsSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_MetricsSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ModelAccessConfigProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelAccessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ModelCardProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelCardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ModelDataQualityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelDataQualityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ModelDataSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ModelMetricsProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelMetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ModelPackageContainerDefinitionProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelPackageContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ModelPackageStatusDetailsProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelPackageStatusDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ModelPackageStatusItemProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelPackageStatusItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ModelQualityProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ModelQualityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.S3DataSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_S3DataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.S3ModelDataSourceProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_S3ModelDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.SecurityConfigProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_SecurityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.SourceAlgorithmProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_SourceAlgorithmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.SourceAlgorithmSpecificationProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_SourceAlgorithmSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.TransformInputProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_TransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.TransformJobDefinitionProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_TransformJobDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.TransformOutputProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_TransformOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.TransformResourcesProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_TransformResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ValidationProfileProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ValidationProfileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPackagePropsMixin.ValidationSpecificationProperty",
		reflect.TypeOf((*CfnModelPackagePropsMixin_ValidationSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin",
		reflect.TypeOf((*CfnModelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin.ContainerDefinitionProperty",
		reflect.TypeOf((*CfnModelPropsMixin_ContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin.HubAccessConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_HubAccessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin.ImageConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_ImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin.InferenceExecutionConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_InferenceExecutionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin.ModelAccessConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_ModelAccessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin.ModelDataSourceProperty",
		reflect.TypeOf((*CfnModelPropsMixin_ModelDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin.MultiModelConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_MultiModelConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin.RepositoryAuthConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_RepositoryAuthConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin.S3DataSourceProperty",
		reflect.TypeOf((*CfnModelPropsMixin_S3DataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnModelPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionMixinProps",
		reflect.TypeOf((*CfnModelQualityJobDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelQualityJobDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.BatchTransformInputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_BatchTransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.ConstraintsResourceProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_ConstraintsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.CsvProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.DatasetFormatProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_DatasetFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.EndpointInputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_EndpointInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.JsonProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_JsonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.ModelQualityAppSpecificationProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_ModelQualityAppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.ModelQualityBaselineConfigProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_ModelQualityBaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.ModelQualityJobInputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_ModelQualityJobInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.MonitoringGroundTruthS3InputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_MonitoringGroundTruthS3InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.MonitoringOutputConfigProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_MonitoringOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.MonitoringOutputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_MonitoringOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.MonitoringResourcesProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_MonitoringResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnModelQualityJobDefinitionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnModelQualityJobDefinitionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringScheduleMixinProps",
		reflect.TypeOf((*CfnMonitoringScheduleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMonitoringSchedulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.BaselineConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_BaselineConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.BatchTransformInputProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_BatchTransformInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.ConstraintsResourceProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_ConstraintsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.CsvProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.DatasetFormatProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_DatasetFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.EndpointInputProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_EndpointInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.JsonProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_JsonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.MonitoringAppSpecificationProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringAppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.MonitoringExecutionSummaryProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringExecutionSummaryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.MonitoringInputProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.MonitoringJobDefinitionProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringJobDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.MonitoringOutputConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.MonitoringOutputProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.MonitoringResourcesProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.MonitoringScheduleConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_MonitoringScheduleConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.ScheduleConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_ScheduleConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.StatisticsResourceProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_StatisticsResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnMonitoringSchedulePropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnMonitoringSchedulePropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnNotebookInstanceLifecycleConfigMixinProps",
		reflect.TypeOf((*CfnNotebookInstanceLifecycleConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnNotebookInstanceLifecycleConfigPropsMixin",
		reflect.TypeOf((*CfnNotebookInstanceLifecycleConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNotebookInstanceLifecycleConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnNotebookInstanceLifecycleConfigPropsMixin.NotebookInstanceLifecycleHookProperty",
		reflect.TypeOf((*CfnNotebookInstanceLifecycleConfigPropsMixin_NotebookInstanceLifecycleHookProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnNotebookInstanceMixinProps",
		reflect.TypeOf((*CfnNotebookInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnNotebookInstancePropsMixin",
		reflect.TypeOf((*CfnNotebookInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNotebookInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnNotebookInstancePropsMixin.InstanceMetadataServiceConfigurationProperty",
		reflect.TypeOf((*CfnNotebookInstancePropsMixin_InstanceMetadataServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnPartnerAppMixinProps",
		reflect.TypeOf((*CfnPartnerAppMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnPartnerAppPropsMixin",
		reflect.TypeOf((*CfnPartnerAppPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPartnerAppPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnPartnerAppPropsMixin.PartnerAppConfigProperty",
		reflect.TypeOf((*CfnPartnerAppPropsMixin_PartnerAppConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnPartnerAppPropsMixin.PartnerAppMaintenanceConfigProperty",
		reflect.TypeOf((*CfnPartnerAppPropsMixin_PartnerAppMaintenanceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnPipelineMixinProps",
		reflect.TypeOf((*CfnPipelineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnPipelinePropsMixin",
		reflect.TypeOf((*CfnPipelinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipelinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnPipelinePropsMixin.ParallelismConfigurationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ParallelismConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnPipelinePropsMixin.PipelineDefinitionProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_PipelineDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnPipelinePropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobMixinProps",
		reflect.TypeOf((*CfnProcessingJobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin",
		reflect.TypeOf((*CfnProcessingJobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProcessingJobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.AppSpecificationProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_AppSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.AthenaDatasetDefinitionProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_AthenaDatasetDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.DatasetDefinitionProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_DatasetDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.ExperimentConfigProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ExperimentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.FeatureStoreOutputProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_FeatureStoreOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.NetworkConfigProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_NetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.ProcessingInputsObjectProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ProcessingInputsObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.ProcessingOutputConfigProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ProcessingOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.ProcessingOutputsObjectProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ProcessingOutputsObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.ProcessingResourcesProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_ProcessingResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.RedshiftDatasetDefinitionProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_RedshiftDatasetDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.S3InputProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_S3InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.S3OutputProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_S3OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.StoppingConditionProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_StoppingConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProcessingJobPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnProcessingJobPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProjectPropsMixin.CfnStackParameterProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_CfnStackParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProjectPropsMixin.CfnTemplateProviderDetailProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_CfnTemplateProviderDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProjectPropsMixin.ProvisioningParameterProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProvisioningParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProjectPropsMixin.ServiceCatalogProvisionedProductDetailsProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ServiceCatalogProvisionedProductDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProjectPropsMixin.ServiceCatalogProvisioningDetailsProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ServiceCatalogProvisioningDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnProjectPropsMixin.TemplateProviderDetailProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_TemplateProviderDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpaceMixinProps",
		reflect.TypeOf((*CfnSpaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin",
		reflect.TypeOf((*CfnSpacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSpacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.CodeRepositoryProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_CodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.CustomFileSystemProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_CustomFileSystemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.CustomImageProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_CustomImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.EFSFileSystemProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_EFSFileSystemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.EbsStorageSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_EbsStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.FSxLustreFileSystemProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_FSxLustreFileSystemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.JupyterServerAppSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_JupyterServerAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.KernelGatewayAppSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_KernelGatewayAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.OwnershipSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_OwnershipSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.ResourceSpecProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_ResourceSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.S3FileSystemProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_S3FileSystemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.SpaceAppLifecycleManagementProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceAppLifecycleManagementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.SpaceCodeEditorAppSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceCodeEditorAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.SpaceIdleSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceIdleSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.SpaceJupyterLabAppSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceJupyterLabAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.SpaceSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.SpaceSharingSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceSharingSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnSpacePropsMixin.SpaceStorageSettingsProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_SpaceStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnStudioLifecycleConfigMixinProps",
		reflect.TypeOf((*CfnStudioLifecycleConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnStudioLifecycleConfigPropsMixin",
		reflect.TypeOf((*CfnStudioLifecycleConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStudioLifecycleConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfileMixinProps",
		reflect.TypeOf((*CfnUserProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin",
		reflect.TypeOf((*CfnUserProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.AppLifecycleManagementProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_AppLifecycleManagementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.CodeEditorAppSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_CodeEditorAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.CodeRepositoryProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_CodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.CustomFileSystemConfigProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_CustomFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.CustomImageProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_CustomImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.CustomPosixUserConfigProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_CustomPosixUserConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.DefaultEbsStorageSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_DefaultEbsStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.DefaultSpaceStorageSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_DefaultSpaceStorageSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.EFSFileSystemConfigProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_EFSFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.FSxLustreFileSystemConfigProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_FSxLustreFileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.HiddenSageMakerImageProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_HiddenSageMakerImageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.IdleSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_IdleSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.JupyterLabAppSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_JupyterLabAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.JupyterServerAppSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_JupyterServerAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.KernelGatewayAppSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_KernelGatewayAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.RStudioServerProAppSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_RStudioServerProAppSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.ResourceSpecProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_ResourceSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.S3FileSystemConfigProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_S3FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.SharingSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_SharingSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.StudioWebPortalSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_StudioWebPortalSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnUserProfilePropsMixin.UserSettingsProperty",
		reflect.TypeOf((*CfnUserProfilePropsMixin_UserSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnWorkteamMixinProps",
		reflect.TypeOf((*CfnWorkteamMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnWorkteamPropsMixin",
		reflect.TypeOf((*CfnWorkteamPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkteamPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnWorkteamPropsMixin.CognitoMemberDefinitionProperty",
		reflect.TypeOf((*CfnWorkteamPropsMixin_CognitoMemberDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnWorkteamPropsMixin.MemberDefinitionProperty",
		reflect.TypeOf((*CfnWorkteamPropsMixin_MemberDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnWorkteamPropsMixin.NotificationConfigurationProperty",
		reflect.TypeOf((*CfnWorkteamPropsMixin_NotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnWorkteamPropsMixin.OidcMemberDefinitionProperty",
		reflect.TypeOf((*CfnWorkteamPropsMixin_OidcMemberDefinitionProperty)(nil)).Elem(),
	)
}
