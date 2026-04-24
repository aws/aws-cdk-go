package awsbatch

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnComputeEnvironmentMixinProps",
		reflect.TypeOf((*CfnComputeEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnComputeEnvironmentPropsMixin",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComputeEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnComputeEnvironmentPropsMixin.ComputeResourcesProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_ComputeResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnComputeEnvironmentPropsMixin.ComputeScalingPolicyProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_ComputeScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnComputeEnvironmentPropsMixin.Ec2ConfigurationObjectProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_Ec2ConfigurationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnComputeEnvironmentPropsMixin.EksConfigurationProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_EksConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnComputeEnvironmentPropsMixin.LaunchTemplateSpecificationOverrideProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_LaunchTemplateSpecificationOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnComputeEnvironmentPropsMixin.LaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_LaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnComputeEnvironmentPropsMixin.UpdatePolicyProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_UpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnConsumableResourceMixinProps",
		reflect.TypeOf((*CfnConsumableResourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnConsumableResourcePropsMixin",
		reflect.TypeOf((*CfnConsumableResourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConsumableResourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionMixinProps",
		reflect.TypeOf((*CfnJobDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJobDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.AuthorizationConfigProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_AuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.ConsumableResourcePropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ConsumableResourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.ConsumableResourceRequirementProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ConsumableResourceRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.ContainerPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ContainerPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.DeviceProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_DeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EcsPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EcsPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EcsTaskPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EcsTaskPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EfsVolumeConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EfsVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EksContainerEnvironmentVariableProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksContainerEnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EksContainerProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksContainerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EksContainerVolumeMountProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksContainerVolumeMountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EksPersistentVolumeClaimProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksPersistentVolumeClaimProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EksPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EksSecretProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksSecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EksVolumeProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksVolumeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EmptyDirProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EmptyDirProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EnvironmentProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EphemeralStorageProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EphemeralStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.EvaluateOnExitProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EvaluateOnExitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.FargatePlatformConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_FargatePlatformConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.FirelensConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_FirelensConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.HostPathProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_HostPathProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.ImagePullSecretProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ImagePullSecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.JobTimeoutProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_JobTimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.LinuxParametersProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_LinuxParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.MetadataProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_MetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.MountPointProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_MountPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.MountPointsProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_MountPointsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.MultiNodeEcsPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_MultiNodeEcsPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.MultiNodeEcsTaskPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_MultiNodeEcsTaskPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.NodePropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_NodePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.NodeRangePropertyProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_NodeRangePropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.PodPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_PodPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.RepositoryCredentialsProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_RepositoryCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.ResourceRequirementProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ResourceRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.ResourceRetentionPolicyProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ResourceRetentionPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.ResourcesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.RetryStrategyProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_RetryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.RuntimePlatformProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_RuntimePlatformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.S3FilesVolumeConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_S3FilesVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.SecretProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.SecurityContextProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_SecurityContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.TaskContainerDependencyProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_TaskContainerDependencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.TaskContainerPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_TaskContainerPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.TimeoutProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_TimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.TmpfsProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_TmpfsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.UlimitProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_UlimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.VolumesHostProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_VolumesHostProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobDefinitionPropsMixin.VolumesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_VolumesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobQueueMixinProps",
		reflect.TypeOf((*CfnJobQueueMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobQueuePropsMixin",
		reflect.TypeOf((*CfnJobQueuePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJobQueuePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobQueuePropsMixin.ComputeEnvironmentOrderProperty",
		reflect.TypeOf((*CfnJobQueuePropsMixin_ComputeEnvironmentOrderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobQueuePropsMixin.JobStateTimeLimitActionProperty",
		reflect.TypeOf((*CfnJobQueuePropsMixin_JobStateTimeLimitActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnJobQueuePropsMixin.ServiceEnvironmentOrderProperty",
		reflect.TypeOf((*CfnJobQueuePropsMixin_ServiceEnvironmentOrderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnQuotaShareMixinProps",
		reflect.TypeOf((*CfnQuotaShareMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnQuotaSharePropsMixin",
		reflect.TypeOf((*CfnQuotaSharePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQuotaSharePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnQuotaSharePropsMixin.QuotaShareCapacityLimitProperty",
		reflect.TypeOf((*CfnQuotaSharePropsMixin_QuotaShareCapacityLimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnQuotaSharePropsMixin.QuotaSharePreemptionConfigurationProperty",
		reflect.TypeOf((*CfnQuotaSharePropsMixin_QuotaSharePreemptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnQuotaSharePropsMixin.QuotaShareResourceSharingConfigurationProperty",
		reflect.TypeOf((*CfnQuotaSharePropsMixin_QuotaShareResourceSharingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnSchedulingPolicyMixinProps",
		reflect.TypeOf((*CfnSchedulingPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnSchedulingPolicyPropsMixin",
		reflect.TypeOf((*CfnSchedulingPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchedulingPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnSchedulingPolicyPropsMixin.FairsharePolicyProperty",
		reflect.TypeOf((*CfnSchedulingPolicyPropsMixin_FairsharePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnSchedulingPolicyPropsMixin.QuotaSharePolicyProperty",
		reflect.TypeOf((*CfnSchedulingPolicyPropsMixin_QuotaSharePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnSchedulingPolicyPropsMixin.ShareAttributesProperty",
		reflect.TypeOf((*CfnSchedulingPolicyPropsMixin_ShareAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnServiceEnvironmentMixinProps",
		reflect.TypeOf((*CfnServiceEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnServiceEnvironmentPropsMixin",
		reflect.TypeOf((*CfnServiceEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServiceEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnServiceEnvironmentPropsMixin.CapacityLimitProperty",
		reflect.TypeOf((*CfnServiceEnvironmentPropsMixin_CapacityLimitProperty)(nil)).Elem(),
	)
}
