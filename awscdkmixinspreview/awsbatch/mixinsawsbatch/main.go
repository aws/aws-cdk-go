package mixinsawsbatch

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentMixinProps",
		reflect.TypeOf((*CfnComputeEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComputeEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin.ComputeResourcesProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_ComputeResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin.Ec2ConfigurationObjectProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_Ec2ConfigurationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin.EksConfigurationProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_EksConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin.LaunchTemplateSpecificationOverrideProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_LaunchTemplateSpecificationOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin.LaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_LaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnComputeEnvironmentPropsMixin.UpdatePolicyProperty",
		reflect.TypeOf((*CfnComputeEnvironmentPropsMixin_UpdatePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnConsumableResourceMixinProps",
		reflect.TypeOf((*CfnConsumableResourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnConsumableResourcePropsMixin",
		reflect.TypeOf((*CfnConsumableResourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConsumableResourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionMixinProps",
		reflect.TypeOf((*CfnJobDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJobDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.AuthorizationConfigProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_AuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.ConsumableResourcePropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ConsumableResourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.ConsumableResourceRequirementProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ConsumableResourceRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.ContainerPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ContainerPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.DeviceProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_DeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EcsPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EcsPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EcsTaskPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EcsTaskPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EfsVolumeConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EfsVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EksContainerEnvironmentVariableProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksContainerEnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EksContainerProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksContainerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EksContainerVolumeMountProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksContainerVolumeMountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EksPersistentVolumeClaimProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksPersistentVolumeClaimProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EksPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EksSecretProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksSecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EksVolumeProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EksVolumeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EmptyDirProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EmptyDirProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EnvironmentProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EphemeralStorageProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EphemeralStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.EvaluateOnExitProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_EvaluateOnExitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.FargatePlatformConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_FargatePlatformConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.FirelensConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_FirelensConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.HostPathProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_HostPathProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.ImagePullSecretProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ImagePullSecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.JobTimeoutProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_JobTimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.LinuxParametersProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_LinuxParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.MetadataProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_MetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.MountPointProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_MountPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.MountPointsProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_MountPointsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.MultiNodeEcsPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_MultiNodeEcsPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.MultiNodeEcsTaskPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_MultiNodeEcsTaskPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.NodePropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_NodePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.NodeRangePropertyProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_NodeRangePropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.PodPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_PodPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.RepositoryCredentialsProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_RepositoryCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.ResourceRequirementProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ResourceRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.ResourceRetentionPolicyProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ResourceRetentionPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.ResourcesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_ResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.RetryStrategyProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_RetryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.RuntimePlatformProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_RuntimePlatformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.SecretProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.SecurityContextProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_SecurityContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.TaskContainerDependencyProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_TaskContainerDependencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.TaskContainerPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_TaskContainerPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.TimeoutProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_TimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.TmpfsProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_TmpfsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.UlimitProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_UlimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.VolumesHostProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_VolumesHostProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobDefinitionPropsMixin.VolumesProperty",
		reflect.TypeOf((*CfnJobDefinitionPropsMixin_VolumesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobQueueMixinProps",
		reflect.TypeOf((*CfnJobQueueMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobQueuePropsMixin",
		reflect.TypeOf((*CfnJobQueuePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJobQueuePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobQueuePropsMixin.ComputeEnvironmentOrderProperty",
		reflect.TypeOf((*CfnJobQueuePropsMixin_ComputeEnvironmentOrderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobQueuePropsMixin.JobStateTimeLimitActionProperty",
		reflect.TypeOf((*CfnJobQueuePropsMixin_JobStateTimeLimitActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobQueuePropsMixin.ServiceEnvironmentOrderProperty",
		reflect.TypeOf((*CfnJobQueuePropsMixin_ServiceEnvironmentOrderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnSchedulingPolicyMixinProps",
		reflect.TypeOf((*CfnSchedulingPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnSchedulingPolicyPropsMixin",
		reflect.TypeOf((*CfnSchedulingPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchedulingPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnSchedulingPolicyPropsMixin.FairsharePolicyProperty",
		reflect.TypeOf((*CfnSchedulingPolicyPropsMixin_FairsharePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnSchedulingPolicyPropsMixin.ShareAttributesProperty",
		reflect.TypeOf((*CfnSchedulingPolicyPropsMixin_ShareAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnServiceEnvironmentMixinProps",
		reflect.TypeOf((*CfnServiceEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnServiceEnvironmentPropsMixin",
		reflect.TypeOf((*CfnServiceEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServiceEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnServiceEnvironmentPropsMixin.CapacityLimitProperty",
		reflect.TypeOf((*CfnServiceEnvironmentPropsMixin_CapacityLimitProperty)(nil)).Elem(),
	)
}
