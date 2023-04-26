package awsautoscaling

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.AdjustmentTier",
		reflect.TypeOf((*AdjustmentTier)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.AdjustmentType",
		reflect.TypeOf((*AdjustmentType)(nil)).Elem(),
		map[string]interface{}{
			"CHANGE_IN_CAPACITY": AdjustmentType_CHANGE_IN_CAPACITY,
			"PERCENT_CHANGE_IN_CAPACITY": AdjustmentType_PERCENT_CHANGE_IN_CAPACITY,
			"EXACT_CAPACITY": AdjustmentType_EXACT_CAPACITY,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.ApplyCloudFormationInitOptions",
		reflect.TypeOf((*ApplyCloudFormationInitOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroup",
		reflect.TypeOf((*AutoScalingGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addLifecycleHook", GoMethod: "AddLifecycleHook"},
			_jsii_.MemberMethod{JsiiMethod: "addSecurityGroup", GoMethod: "AddSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addUserData", GoMethod: "AddUserData"},
			_jsii_.MemberMethod{JsiiMethod: "addWarmPool", GoMethod: "AddWarmPool"},
			_jsii_.MemberProperty{JsiiProperty: "albTargetGroup", GoGetter: "AlbTargetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "applyCloudFormationInit", GoMethod: "ApplyCloudFormationInit"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "areNewInstancesProtectedFromScaleIn", GoMethod: "AreNewInstancesProtectedFromScaleIn"},
			_jsii_.MemberMethod{JsiiMethod: "attachToApplicationTargetGroup", GoMethod: "AttachToApplicationTargetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "attachToClassicLB", GoMethod: "AttachToClassicLB"},
			_jsii_.MemberMethod{JsiiMethod: "attachToNetworkTargetGroup", GoMethod: "AttachToNetworkTargetGroup"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupArn", GoGetter: "AutoScalingGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupName", GoGetter: "AutoScalingGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "hasCalledScaleOnRequestCount", GoGetter: "HasCalledScaleOnRequestCount"},
			_jsii_.MemberProperty{JsiiProperty: "maxInstanceLifetime", GoGetter: "MaxInstanceLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "newInstancesProtectedFromScaleIn", GoGetter: "NewInstancesProtectedFromScaleIn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "osType", GoGetter: "OsType"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "protectNewInstancesFromScaleIn", GoMethod: "ProtectNewInstancesFromScaleIn"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnCpuUtilization", GoMethod: "ScaleOnCpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnIncomingBytes", GoMethod: "ScaleOnIncomingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnMetric", GoMethod: "ScaleOnMetric"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnOutgoingBytes", GoMethod: "ScaleOnOutgoingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnRequestCount", GoMethod: "ScaleOnRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnSchedule", GoMethod: "ScaleOnSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "scaleToTrackMetric", GoMethod: "ScaleToTrackMetric"},
			_jsii_.MemberProperty{JsiiProperty: "spotPrice", GoGetter: "SpotPrice"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userData", GoGetter: "UserData"},
		},
		func() interface{} {
			j := jsiiProxy_AutoScalingGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAutoScalingGroup)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awselasticloadbalancingILoadBalancerTarget)
			_jsii_.InitJsiiProxy(&j.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget)
			_jsii_.InitJsiiProxy(&j.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroupProps",
		reflect.TypeOf((*AutoScalingGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroupRequireImdsv2Aspect",
		reflect.TypeOf((*AutoScalingGroupRequireImdsv2Aspect)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
			_jsii_.MemberMethod{JsiiMethod: "warn", GoMethod: "Warn"},
		},
		func() interface{} {
			j := jsiiProxy_AutoScalingGroupRequireImdsv2Aspect{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.BaseTargetTrackingProps",
		reflect.TypeOf((*BaseTargetTrackingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.BasicLifecycleHookProps",
		reflect.TypeOf((*BasicLifecycleHookProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.BasicScheduledActionProps",
		reflect.TypeOf((*BasicScheduledActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.BasicStepScalingPolicyProps",
		reflect.TypeOf((*BasicStepScalingPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.BasicTargetTrackingScalingPolicyProps",
		reflect.TypeOf((*BasicTargetTrackingScalingPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.BindHookTargetOptions",
		reflect.TypeOf((*BindHookTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.BlockDevice",
		reflect.TypeOf((*BlockDevice)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		reflect.TypeOf((*BlockDeviceVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ebsDevice", GoGetter: "EbsDevice"},
			_jsii_.MemberProperty{JsiiProperty: "virtualName", GoGetter: "VirtualName"},
		},
		func() interface{} {
			return &jsiiProxy_BlockDeviceVolume{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		reflect.TypeOf((*CfnAutoScalingGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupName", GoGetter: "AutoScalingGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "capacityRebalance", GoGetter: "CapacityRebalance"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "cooldown", GoGetter: "Cooldown"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInstanceWarmup", GoGetter: "DefaultInstanceWarmup"},
			_jsii_.MemberProperty{JsiiProperty: "desiredCapacity", GoGetter: "DesiredCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "desiredCapacityType", GoGetter: "DesiredCapacityType"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckGracePeriod", GoGetter: "HealthCheckGracePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckType", GoGetter: "HealthCheckType"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "launchConfigurationName", GoGetter: "LaunchConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplate", GoGetter: "LaunchTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleHookSpecificationList", GoGetter: "LifecycleHookSpecificationList"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerNames", GoGetter: "LoadBalancerNames"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maxInstanceLifetime", GoGetter: "MaxInstanceLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "maxSize", GoGetter: "MaxSize"},
			_jsii_.MemberProperty{JsiiProperty: "metricsCollection", GoGetter: "MetricsCollection"},
			_jsii_.MemberProperty{JsiiProperty: "minSize", GoGetter: "MinSize"},
			_jsii_.MemberProperty{JsiiProperty: "mixedInstancesPolicy", GoGetter: "MixedInstancesPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "newInstancesProtectedFromScaleIn", GoGetter: "NewInstancesProtectedFromScaleIn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationConfigurations", GoGetter: "NotificationConfigurations"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "placementGroup", GoGetter: "PlacementGroup"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "serviceLinkedRoleArn", GoGetter: "ServiceLinkedRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupArns", GoGetter: "TargetGroupArns"},
			_jsii_.MemberProperty{JsiiProperty: "terminationPolicies", GoGetter: "TerminationPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcZoneIdentifier", GoGetter: "VpcZoneIdentifier"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAutoScalingGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.AcceleratorCountRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_AcceleratorCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.AcceleratorTotalMemoryMiBRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_AcceleratorTotalMemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.BaselineEbsBandwidthMbpsRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_BaselineEbsBandwidthMbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.InstanceRequirementsProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_InstanceRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.InstancesDistributionProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_InstancesDistributionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.LaunchTemplateOverridesProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_LaunchTemplateOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.LaunchTemplateProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_LaunchTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.LaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_LaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.LifecycleHookSpecificationProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_LifecycleHookSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.MemoryGiBPerVCpuRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_MemoryGiBPerVCpuRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.MemoryMiBRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_MemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.MetricsCollectionProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_MetricsCollectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.MixedInstancesPolicyProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_MixedInstancesPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.NetworkBandwidthGbpsRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_NetworkBandwidthGbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.NetworkInterfaceCountRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_NetworkInterfaceCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.NotificationConfigurationProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_NotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.TagPropertyProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_TagPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.TotalLocalStorageGBRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_TotalLocalStorageGBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup.VCpuCountRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroup_VCpuCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroupProps",
		reflect.TypeOf((*CfnAutoScalingGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfiguration",
		reflect.TypeOf((*CfnLaunchConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "associatePublicIpAddress", GoGetter: "AssociatePublicIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "blockDeviceMappings", GoGetter: "BlockDeviceMappings"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "classicLinkVpcId", GoGetter: "ClassicLinkVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "classicLinkVpcSecurityGroups", GoGetter: "ClassicLinkVpcSecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimized", GoGetter: "EbsOptimized"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "iamInstanceProfile", GoGetter: "IamInstanceProfile"},
			_jsii_.MemberProperty{JsiiProperty: "imageId", GoGetter: "ImageId"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "instanceMonitoring", GoGetter: "InstanceMonitoring"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "kernelId", GoGetter: "KernelId"},
			_jsii_.MemberProperty{JsiiProperty: "keyName", GoGetter: "KeyName"},
			_jsii_.MemberProperty{JsiiProperty: "launchConfigurationName", GoGetter: "LaunchConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "metadataOptions", GoGetter: "MetadataOptions"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "placementTenancy", GoGetter: "PlacementTenancy"},
			_jsii_.MemberProperty{JsiiProperty: "ramDiskId", GoGetter: "RamDiskId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spotPrice", GoGetter: "SpotPrice"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberProperty{JsiiProperty: "userData", GoGetter: "UserData"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLaunchConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfiguration.BlockDeviceMappingProperty",
		reflect.TypeOf((*CfnLaunchConfiguration_BlockDeviceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfiguration.BlockDeviceProperty",
		reflect.TypeOf((*CfnLaunchConfiguration_BlockDeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfiguration.MetadataOptionsProperty",
		reflect.TypeOf((*CfnLaunchConfiguration_MetadataOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfigurationProps",
		reflect.TypeOf((*CfnLaunchConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.CfnLifecycleHook",
		reflect.TypeOf((*CfnLifecycleHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupName", GoGetter: "AutoScalingGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResult", GoGetter: "DefaultResult"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "heartbeatTimeout", GoGetter: "HeartbeatTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleHookName", GoGetter: "LifecycleHookName"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleTransition", GoGetter: "LifecycleTransition"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationMetadata", GoGetter: "NotificationMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "notificationTargetArn", GoGetter: "NotificationTargetArn"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLifecycleHook{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnLifecycleHookProps",
		reflect.TypeOf((*CfnLifecycleHookProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		reflect.TypeOf((*CfnScalingPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "adjustmentType", GoGetter: "AdjustmentType"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrPolicyName", GoGetter: "AttrPolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupName", GoGetter: "AutoScalingGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cooldown", GoGetter: "Cooldown"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "estimatedInstanceWarmup", GoGetter: "EstimatedInstanceWarmup"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "metricAggregationType", GoGetter: "MetricAggregationType"},
			_jsii_.MemberProperty{JsiiProperty: "minAdjustmentMagnitude", GoGetter: "MinAdjustmentMagnitude"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyType", GoGetter: "PolicyType"},
			_jsii_.MemberProperty{JsiiProperty: "predictiveScalingConfiguration", GoGetter: "PredictiveScalingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "scalingAdjustment", GoGetter: "ScalingAdjustment"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stepAdjustments", GoGetter: "StepAdjustments"},
			_jsii_.MemberProperty{JsiiProperty: "targetTrackingConfiguration", GoGetter: "TargetTrackingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScalingPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.CustomizedMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPolicy_CustomizedMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.MetricDataQueryProperty",
		reflect.TypeOf((*CfnScalingPolicy_MetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.MetricDimensionProperty",
		reflect.TypeOf((*CfnScalingPolicy_MetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.MetricProperty",
		reflect.TypeOf((*CfnScalingPolicy_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.MetricStatProperty",
		reflect.TypeOf((*CfnScalingPolicy_MetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.PredefinedMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPolicy_PredefinedMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.PredictiveScalingConfigurationProperty",
		reflect.TypeOf((*CfnScalingPolicy_PredictiveScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.PredictiveScalingCustomizedCapacityMetricProperty",
		reflect.TypeOf((*CfnScalingPolicy_PredictiveScalingCustomizedCapacityMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.PredictiveScalingCustomizedLoadMetricProperty",
		reflect.TypeOf((*CfnScalingPolicy_PredictiveScalingCustomizedLoadMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.PredictiveScalingCustomizedScalingMetricProperty",
		reflect.TypeOf((*CfnScalingPolicy_PredictiveScalingCustomizedScalingMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.PredictiveScalingMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPolicy_PredictiveScalingMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.PredictiveScalingPredefinedLoadMetricProperty",
		reflect.TypeOf((*CfnScalingPolicy_PredictiveScalingPredefinedLoadMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.PredictiveScalingPredefinedMetricPairProperty",
		reflect.TypeOf((*CfnScalingPolicy_PredictiveScalingPredefinedMetricPairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.PredictiveScalingPredefinedScalingMetricProperty",
		reflect.TypeOf((*CfnScalingPolicy_PredictiveScalingPredefinedScalingMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.StepAdjustmentProperty",
		reflect.TypeOf((*CfnScalingPolicy_StepAdjustmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy.TargetTrackingConfigurationProperty",
		reflect.TypeOf((*CfnScalingPolicy_TargetTrackingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicyProps",
		reflect.TypeOf((*CfnScalingPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.CfnScheduledAction",
		reflect.TypeOf((*CfnScheduledAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrScheduledActionName", GoGetter: "AttrScheduledActionName"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupName", GoGetter: "AutoScalingGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "desiredCapacity", GoGetter: "DesiredCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "endTime", GoGetter: "EndTime"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maxSize", GoGetter: "MaxSize"},
			_jsii_.MemberProperty{JsiiProperty: "minSize", GoGetter: "MinSize"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "recurrence", GoGetter: "Recurrence"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "startTime", GoGetter: "StartTime"},
			_jsii_.MemberProperty{JsiiProperty: "timeZone", GoGetter: "TimeZone"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScheduledAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnScheduledActionProps",
		reflect.TypeOf((*CfnScheduledActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.CfnWarmPool",
		reflect.TypeOf((*CfnWarmPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupName", GoGetter: "AutoScalingGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceReusePolicy", GoGetter: "InstanceReusePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maxGroupPreparedCapacity", GoGetter: "MaxGroupPreparedCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "minSize", GoGetter: "MinSize"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "poolState", GoGetter: "PoolState"},
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
			j := jsiiProxy_CfnWarmPool{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnWarmPool.InstanceReusePolicyProperty",
		reflect.TypeOf((*CfnWarmPool_InstanceReusePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CfnWarmPoolProps",
		reflect.TypeOf((*CfnWarmPoolProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CommonAutoScalingGroupProps",
		reflect.TypeOf((*CommonAutoScalingGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CpuUtilizationScalingProps",
		reflect.TypeOf((*CpuUtilizationScalingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.CronOptions",
		reflect.TypeOf((*CronOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.DefaultResult",
		reflect.TypeOf((*DefaultResult)(nil)).Elem(),
		map[string]interface{}{
			"CONTINUE": DefaultResult_CONTINUE,
			"ABANDON": DefaultResult_ABANDON,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.EbsDeviceOptions",
		reflect.TypeOf((*EbsDeviceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.EbsDeviceOptionsBase",
		reflect.TypeOf((*EbsDeviceOptionsBase)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.EbsDeviceProps",
		reflect.TypeOf((*EbsDeviceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.EbsDeviceSnapshotOptions",
		reflect.TypeOf((*EbsDeviceSnapshotOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.EbsDeviceVolumeType",
		reflect.TypeOf((*EbsDeviceVolumeType)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": EbsDeviceVolumeType_STANDARD,
			"IO1": EbsDeviceVolumeType_IO1,
			"GP2": EbsDeviceVolumeType_GP2,
			"GP3": EbsDeviceVolumeType_GP3,
			"ST1": EbsDeviceVolumeType_ST1,
			"SC1": EbsDeviceVolumeType_SC1,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.Ec2HealthCheckOptions",
		reflect.TypeOf((*Ec2HealthCheckOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.ElbHealthCheckOptions",
		reflect.TypeOf((*ElbHealthCheckOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		reflect.TypeOf((*GroupMetric)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_GroupMetric{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.GroupMetrics",
		reflect.TypeOf((*GroupMetrics)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GroupMetrics{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.HealthCheck",
		reflect.TypeOf((*HealthCheck)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "gracePeriod", GoGetter: "GracePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_HealthCheck{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_autoscaling.IAutoScalingGroup",
		reflect.TypeOf((*IAutoScalingGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addLifecycleHook", GoMethod: "AddLifecycleHook"},
			_jsii_.MemberMethod{JsiiMethod: "addUserData", GoMethod: "AddUserData"},
			_jsii_.MemberMethod{JsiiMethod: "addWarmPool", GoMethod: "AddWarmPool"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupArn", GoGetter: "AutoScalingGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupName", GoGetter: "AutoScalingGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "osType", GoGetter: "OsType"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnCpuUtilization", GoMethod: "ScaleOnCpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnIncomingBytes", GoMethod: "ScaleOnIncomingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnMetric", GoMethod: "ScaleOnMetric"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnOutgoingBytes", GoMethod: "ScaleOnOutgoingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnSchedule", GoMethod: "ScaleOnSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "scaleToTrackMetric", GoMethod: "ScaleToTrackMetric"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAutoScalingGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_autoscaling.ILifecycleHook",
		reflect.TypeOf((*ILifecycleHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ILifecycleHook{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_autoscaling.ILifecycleHookTarget",
		reflect.TypeOf((*ILifecycleHookTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ILifecycleHookTarget{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.InstancesDistribution",
		reflect.TypeOf((*InstancesDistribution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.LaunchTemplateOverrides",
		reflect.TypeOf((*LaunchTemplateOverrides)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.LifecycleHook",
		reflect.TypeOf((*LifecycleHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleHookName", GoGetter: "LifecycleHookName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LifecycleHook{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILifecycleHook)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.LifecycleHookProps",
		reflect.TypeOf((*LifecycleHookProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.LifecycleHookTargetConfig",
		reflect.TypeOf((*LifecycleHookTargetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.LifecycleTransition",
		reflect.TypeOf((*LifecycleTransition)(nil)).Elem(),
		map[string]interface{}{
			"INSTANCE_LAUNCHING": LifecycleTransition_INSTANCE_LAUNCHING,
			"INSTANCE_TERMINATING": LifecycleTransition_INSTANCE_TERMINATING,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.MetricAggregationType",
		reflect.TypeOf((*MetricAggregationType)(nil)).Elem(),
		map[string]interface{}{
			"AVERAGE": MetricAggregationType_AVERAGE,
			"MINIMUM": MetricAggregationType_MINIMUM,
			"MAXIMUM": MetricAggregationType_MAXIMUM,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.MetricTargetTrackingProps",
		reflect.TypeOf((*MetricTargetTrackingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.MixedInstancesPolicy",
		reflect.TypeOf((*MixedInstancesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.Monitoring",
		reflect.TypeOf((*Monitoring)(nil)).Elem(),
		map[string]interface{}{
			"BASIC": Monitoring_BASIC,
			"DETAILED": Monitoring_DETAILED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.NetworkUtilizationScalingProps",
		reflect.TypeOf((*NetworkUtilizationScalingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.NotificationConfiguration",
		reflect.TypeOf((*NotificationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.OnDemandAllocationStrategy",
		reflect.TypeOf((*OnDemandAllocationStrategy)(nil)).Elem(),
		map[string]interface{}{
			"PRIORITIZED": OnDemandAllocationStrategy_PRIORITIZED,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.PoolState",
		reflect.TypeOf((*PoolState)(nil)).Elem(),
		map[string]interface{}{
			"HIBERNATED": PoolState_HIBERNATED,
			"RUNNING": PoolState_RUNNING,
			"STOPPED": PoolState_STOPPED,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.PredefinedMetric",
		reflect.TypeOf((*PredefinedMetric)(nil)).Elem(),
		map[string]interface{}{
			"ASG_AVERAGE_CPU_UTILIZATION": PredefinedMetric_ASG_AVERAGE_CPU_UTILIZATION,
			"ASG_AVERAGE_NETWORK_IN": PredefinedMetric_ASG_AVERAGE_NETWORK_IN,
			"ASG_AVERAGE_NETWORK_OUT": PredefinedMetric_ASG_AVERAGE_NETWORK_OUT,
			"ALB_REQUEST_COUNT_PER_TARGET": PredefinedMetric_ALB_REQUEST_COUNT_PER_TARGET,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.RenderSignalsOptions",
		reflect.TypeOf((*RenderSignalsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.RequestCountScalingProps",
		reflect.TypeOf((*RequestCountScalingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.RollingUpdateOptions",
		reflect.TypeOf((*RollingUpdateOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.ScalingEvent",
		reflect.TypeOf((*ScalingEvent)(nil)).Elem(),
		map[string]interface{}{
			"INSTANCE_LAUNCH": ScalingEvent_INSTANCE_LAUNCH,
			"INSTANCE_TERMINATE": ScalingEvent_INSTANCE_TERMINATE,
			"INSTANCE_TERMINATE_ERROR": ScalingEvent_INSTANCE_TERMINATE_ERROR,
			"INSTANCE_LAUNCH_ERROR": ScalingEvent_INSTANCE_LAUNCH_ERROR,
			"TEST_NOTIFICATION": ScalingEvent_TEST_NOTIFICATION,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.ScalingEvents",
		reflect.TypeOf((*ScalingEvents)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ScalingEvents{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.ScalingInterval",
		reflect.TypeOf((*ScalingInterval)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.ScalingProcess",
		reflect.TypeOf((*ScalingProcess)(nil)).Elem(),
		map[string]interface{}{
			"LAUNCH": ScalingProcess_LAUNCH,
			"TERMINATE": ScalingProcess_TERMINATE,
			"HEALTH_CHECK": ScalingProcess_HEALTH_CHECK,
			"REPLACE_UNHEALTHY": ScalingProcess_REPLACE_UNHEALTHY,
			"AZ_REBALANCE": ScalingProcess_AZ_REBALANCE,
			"ALARM_NOTIFICATION": ScalingProcess_ALARM_NOTIFICATION,
			"SCHEDULED_ACTIONS": ScalingProcess_SCHEDULED_ACTIONS,
			"ADD_TO_LOAD_BALANCER": ScalingProcess_ADD_TO_LOAD_BALANCER,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.Schedule",
		reflect.TypeOf((*Schedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "expressionString", GoGetter: "ExpressionString"},
		},
		func() interface{} {
			return &jsiiProxy_Schedule{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.ScheduledAction",
		reflect.TypeOf((*ScheduledAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledActionName", GoGetter: "ScheduledActionName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ScheduledAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.ScheduledActionProps",
		reflect.TypeOf((*ScheduledActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.Signals",
		reflect.TypeOf((*Signals)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "doRender", GoMethod: "DoRender"},
			_jsii_.MemberMethod{JsiiMethod: "renderCreationPolicy", GoMethod: "RenderCreationPolicy"},
		},
		func() interface{} {
			return &jsiiProxy_Signals{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.SignalsOptions",
		reflect.TypeOf((*SignalsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.SpotAllocationStrategy",
		reflect.TypeOf((*SpotAllocationStrategy)(nil)).Elem(),
		map[string]interface{}{
			"LOWEST_PRICE": SpotAllocationStrategy_LOWEST_PRICE,
			"CAPACITY_OPTIMIZED": SpotAllocationStrategy_CAPACITY_OPTIMIZED,
			"CAPACITY_OPTIMIZED_PRIORITIZED": SpotAllocationStrategy_CAPACITY_OPTIMIZED_PRIORITIZED,
			"PRICE_CAPACITY_OPTIMIZED": SpotAllocationStrategy_PRICE_CAPACITY_OPTIMIZED,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.StepScalingAction",
		reflect.TypeOf((*StepScalingAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAdjustment", GoMethod: "AddAdjustment"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "scalingPolicyArn", GoGetter: "ScalingPolicyArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StepScalingAction{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.StepScalingActionProps",
		reflect.TypeOf((*StepScalingActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.StepScalingPolicy",
		reflect.TypeOf((*StepScalingPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "lowerAction", GoGetter: "LowerAction"},
			_jsii_.MemberProperty{JsiiProperty: "lowerAlarm", GoGetter: "LowerAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "upperAction", GoGetter: "UpperAction"},
			_jsii_.MemberProperty{JsiiProperty: "upperAlarm", GoGetter: "UpperAlarm"},
		},
		func() interface{} {
			j := jsiiProxy_StepScalingPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.StepScalingPolicyProps",
		reflect.TypeOf((*StepScalingPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.TargetTrackingScalingPolicy",
		reflect.TypeOf((*TargetTrackingScalingPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "scalingPolicyArn", GoGetter: "ScalingPolicyArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TargetTrackingScalingPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.TargetTrackingScalingPolicyProps",
		reflect.TypeOf((*TargetTrackingScalingPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_autoscaling.TerminationPolicy",
		reflect.TypeOf((*TerminationPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALLOCATION_STRATEGY": TerminationPolicy_ALLOCATION_STRATEGY,
			"CLOSEST_TO_NEXT_INSTANCE_HOUR": TerminationPolicy_CLOSEST_TO_NEXT_INSTANCE_HOUR,
			"DEFAULT": TerminationPolicy_DEFAULT,
			"NEWEST_INSTANCE": TerminationPolicy_NEWEST_INSTANCE,
			"OLDEST_INSTANCE": TerminationPolicy_OLDEST_INSTANCE,
			"OLDEST_LAUNCH_CONFIGURATION": TerminationPolicy_OLDEST_LAUNCH_CONFIGURATION,
			"OLDEST_LAUNCH_TEMPLATE": TerminationPolicy_OLDEST_LAUNCH_TEMPLATE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.UpdatePolicy",
		reflect.TypeOf((*UpdatePolicy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_UpdatePolicy{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling.WarmPool",
		reflect.TypeOf((*WarmPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WarmPool{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.WarmPoolOptions",
		reflect.TypeOf((*WarmPoolOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling.WarmPoolProps",
		reflect.TypeOf((*WarmPoolProps)(nil)).Elem(),
	)
}
