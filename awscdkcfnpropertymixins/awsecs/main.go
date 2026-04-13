package awsecs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderMixinProps",
		reflect.TypeOf((*CfnCapacityProviderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapacityProviderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.AcceleratorCountRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_AcceleratorCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.AcceleratorTotalMemoryMiBRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_AcceleratorTotalMemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.AutoScalingGroupProviderProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_AutoScalingGroupProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.BaselineEbsBandwidthMbpsRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_BaselineEbsBandwidthMbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.CapacityReservationRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_CapacityReservationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.InfrastructureOptimizationProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_InfrastructureOptimizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.InstanceLaunchTemplateProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_InstanceLaunchTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.InstanceRequirementsRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_InstanceRequirementsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.ManagedInstancesLocalStorageConfigurationProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_ManagedInstancesLocalStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.ManagedInstancesNetworkConfigurationProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_ManagedInstancesNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.ManagedInstancesProviderProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_ManagedInstancesProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.ManagedInstancesStorageConfigurationProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_ManagedInstancesStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.ManagedScalingProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_ManagedScalingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.MemoryGiBPerVCpuRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_MemoryGiBPerVCpuRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.MemoryMiBRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_MemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.NetworkBandwidthGbpsRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_NetworkBandwidthGbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.NetworkInterfaceCountRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_NetworkInterfaceCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.TotalLocalStorageGBRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_TotalLocalStorageGBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnCapacityProviderPropsMixin.VCpuCountRangeRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_VCpuCountRangeRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterCapacityProviderAssociationsMixinProps",
		reflect.TypeOf((*CfnClusterCapacityProviderAssociationsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterCapacityProviderAssociationsPropsMixin",
		reflect.TypeOf((*CfnClusterCapacityProviderAssociationsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterCapacityProviderAssociationsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterCapacityProviderAssociationsPropsMixin.CapacityProviderStrategyProperty",
		reflect.TypeOf((*CfnClusterCapacityProviderAssociationsPropsMixin_CapacityProviderStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin.CapacityProviderStrategyItemProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_CapacityProviderStrategyItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin.ClusterConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin.ClusterSettingsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin.ExecuteCommandConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ExecuteCommandConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin.ExecuteCommandLogConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ExecuteCommandLogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin.ManagedStorageConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ManagedStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin.ServiceConnectDefaultsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ServiceConnectDefaultsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonMixinProps",
		reflect.TypeOf((*CfnDaemonMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonPropsMixin",
		reflect.TypeOf((*CfnDaemonPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDaemonPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonPropsMixin.DaemonAlarmConfigurationProperty",
		reflect.TypeOf((*CfnDaemonPropsMixin_DaemonAlarmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonPropsMixin.DaemonDeploymentConfigurationProperty",
		reflect.TypeOf((*CfnDaemonPropsMixin_DaemonDeploymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionMixinProps",
		reflect.TypeOf((*CfnDaemonTaskDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDaemonTaskDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.ContainerDependencyProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_ContainerDependencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.DaemonContainerDefinitionProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_DaemonContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.DeviceProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_DeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.EnvironmentFileProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_EnvironmentFileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.FirelensConfigurationProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_FirelensConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.HealthCheckProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_HealthCheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.HostVolumePropertiesProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_HostVolumePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.KernelCapabilitiesProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_KernelCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.KeyValuePairProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.LinuxParametersProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_LinuxParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.MountPointProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_MountPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.RepositoryCredentialsProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_RepositoryCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.RestartPolicyProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_RestartPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.SecretProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.SystemControlProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_SystemControlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.TmpfsProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_TmpfsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.UlimitProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_UlimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin.VolumeProperty",
		reflect.TypeOf((*CfnDaemonTaskDefinitionPropsMixin_VolumeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServiceMixinProps",
		reflect.TypeOf((*CfnExpressGatewayServiceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExpressGatewayServicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.AutoScalingArnsProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_AutoScalingArnsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.ECSManagedResourceArnsProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ECSManagedResourceArnsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.ExpressGatewayContainerProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayContainerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.ExpressGatewayRepositoryCredentialsProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayRepositoryCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.ExpressGatewayScalingTargetProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayScalingTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.ExpressGatewayServiceAwsLogsConfigurationProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayServiceAwsLogsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.ExpressGatewayServiceConfigurationProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.ExpressGatewayServiceNetworkConfigurationProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayServiceNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.ExpressGatewayServiceStatusProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayServiceStatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.IngressPathArnsProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_IngressPathArnsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.IngressPathSummaryProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_IngressPathSummaryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.KeyValuePairProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnExpressGatewayServicePropsMixin.SecretProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnPrimaryTaskSetMixinProps",
		reflect.TypeOf((*CfnPrimaryTaskSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnPrimaryTaskSetPropsMixin",
		reflect.TypeOf((*CfnPrimaryTaskSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrimaryTaskSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServiceMixinProps",
		reflect.TypeOf((*CfnServiceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin",
		reflect.TypeOf((*CfnServicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.AdvancedConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AdvancedConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.AwsVpcConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AwsVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.CanaryConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CanaryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.CapacityProviderStrategyItemProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CapacityProviderStrategyItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.DeploymentAlarmsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DeploymentAlarmsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.DeploymentCircuitBreakerProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DeploymentCircuitBreakerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.DeploymentConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DeploymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.DeploymentControllerProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DeploymentControllerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.DeploymentLifecycleHookProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DeploymentLifecycleHookProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.EBSTagSpecificationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_EBSTagSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ForceNewDeploymentProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ForceNewDeploymentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.LinearConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_LinearConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.LoadBalancerProperty",
		reflect.TypeOf((*CfnServicePropsMixin_LoadBalancerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.PlacementConstraintProperty",
		reflect.TypeOf((*CfnServicePropsMixin_PlacementConstraintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.PlacementStrategyProperty",
		reflect.TypeOf((*CfnServicePropsMixin_PlacementStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.SecretProperty",
		reflect.TypeOf((*CfnServicePropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceConnectAccessLogConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectAccessLogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceConnectClientAliasProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectClientAliasProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceConnectConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceConnectServiceProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectServiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceConnectTestTrafficRulesHeaderProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectTestTrafficRulesHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceConnectTestTrafficRulesHeaderValueProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectTestTrafficRulesHeaderValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceConnectTestTrafficRulesProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectTestTrafficRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceConnectTlsCertificateAuthorityProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectTlsCertificateAuthorityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceConnectTlsConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectTlsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceManagedEBSVolumeConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceManagedEBSVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceRegistryProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceRegistryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.ServiceVolumeConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.TimeoutConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_TimeoutConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnServicePropsMixin.VpcLatticeConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_VpcLatticeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionMixinProps",
		reflect.TypeOf((*CfnTaskDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTaskDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.AuthorizationConfigProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_AuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.ContainerDefinitionProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_ContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.ContainerDependencyProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_ContainerDependencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.DeviceProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_DeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.DockerVolumeConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_DockerVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.EFSVolumeConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_EFSVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.EnvironmentFileProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_EnvironmentFileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.EphemeralStorageProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_EphemeralStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.FSxAuthorizationConfigProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_FSxAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.FSxWindowsFileServerVolumeConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_FSxWindowsFileServerVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.FirelensConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_FirelensConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.HealthCheckProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_HealthCheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.HostEntryProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_HostEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.HostVolumePropertiesProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_HostVolumePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.InferenceAcceleratorProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_InferenceAcceleratorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.KernelCapabilitiesProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_KernelCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.KeyValuePairProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.LinuxParametersProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_LinuxParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.MountPointProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_MountPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.PortMappingProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_PortMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.ProxyConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_ProxyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.RepositoryCredentialsProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_RepositoryCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.ResourceRequirementProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_ResourceRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.RestartPolicyProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_RestartPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.RuntimePlatformProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_RuntimePlatformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.SecretProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.SystemControlProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_SystemControlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.TaskDefinitionPlacementConstraintProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_TaskDefinitionPlacementConstraintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.TmpfsProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_TmpfsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.UlimitProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_UlimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.VolumeFromProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_VolumeFromProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskDefinitionPropsMixin.VolumeProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_VolumeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskSetMixinProps",
		reflect.TypeOf((*CfnTaskSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskSetPropsMixin",
		reflect.TypeOf((*CfnTaskSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTaskSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskSetPropsMixin.AwsVpcConfigurationProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_AwsVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskSetPropsMixin.CapacityProviderStrategyItemProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_CapacityProviderStrategyItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskSetPropsMixin.LoadBalancerProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_LoadBalancerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskSetPropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskSetPropsMixin.ScaleProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_ScaleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnTaskSetPropsMixin.ServiceRegistryProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_ServiceRegistryProperty)(nil)).Elem(),
	)
}
