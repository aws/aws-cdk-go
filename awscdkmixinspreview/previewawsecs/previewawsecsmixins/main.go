package previewawsecsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderMixinProps",
		reflect.TypeOf((*CfnCapacityProviderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapacityProviderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.AcceleratorCountRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_AcceleratorCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.AcceleratorTotalMemoryMiBRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_AcceleratorTotalMemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.AutoScalingGroupProviderProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_AutoScalingGroupProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.BaselineEbsBandwidthMbpsRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_BaselineEbsBandwidthMbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.InfrastructureOptimizationProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_InfrastructureOptimizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.InstanceLaunchTemplateProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_InstanceLaunchTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.InstanceRequirementsRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_InstanceRequirementsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.ManagedInstancesNetworkConfigurationProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_ManagedInstancesNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.ManagedInstancesProviderProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_ManagedInstancesProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.ManagedInstancesStorageConfigurationProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_ManagedInstancesStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.ManagedScalingProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_ManagedScalingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.MemoryGiBPerVCpuRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_MemoryGiBPerVCpuRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.MemoryMiBRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_MemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.NetworkBandwidthGbpsRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_NetworkBandwidthGbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.NetworkInterfaceCountRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_NetworkInterfaceCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.TotalLocalStorageGBRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_TotalLocalStorageGBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnCapacityProviderPropsMixin.VCpuCountRangeRequestProperty",
		reflect.TypeOf((*CfnCapacityProviderPropsMixin_VCpuCountRangeRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterCapacityProviderAssociationsMixinProps",
		reflect.TypeOf((*CfnClusterCapacityProviderAssociationsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterCapacityProviderAssociationsPropsMixin",
		reflect.TypeOf((*CfnClusterCapacityProviderAssociationsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterCapacityProviderAssociationsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterCapacityProviderAssociationsPropsMixin.CapacityProviderStrategyProperty",
		reflect.TypeOf((*CfnClusterCapacityProviderAssociationsPropsMixin_CapacityProviderStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterPropsMixin.CapacityProviderStrategyItemProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_CapacityProviderStrategyItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterPropsMixin.ClusterConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterPropsMixin.ClusterSettingsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterPropsMixin.ExecuteCommandConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ExecuteCommandConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterPropsMixin.ExecuteCommandLogConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ExecuteCommandLogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterPropsMixin.ManagedStorageConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ManagedStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterPropsMixin.ServiceConnectDefaultsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ServiceConnectDefaultsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServiceMixinProps",
		reflect.TypeOf((*CfnExpressGatewayServiceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExpressGatewayServicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin.ExpressGatewayContainerProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayContainerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin.ExpressGatewayRepositoryCredentialsProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayRepositoryCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin.ExpressGatewayScalingTargetProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayScalingTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin.ExpressGatewayServiceAwsLogsConfigurationProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayServiceAwsLogsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin.ExpressGatewayServiceConfigurationProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin.ExpressGatewayServiceNetworkConfigurationProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayServiceNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin.ExpressGatewayServiceStatusProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_ExpressGatewayServiceStatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin.IngressPathSummaryProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_IngressPathSummaryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin.KeyValuePairProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin.SecretProperty",
		reflect.TypeOf((*CfnExpressGatewayServicePropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnPrimaryTaskSetMixinProps",
		reflect.TypeOf((*CfnPrimaryTaskSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnPrimaryTaskSetPropsMixin",
		reflect.TypeOf((*CfnPrimaryTaskSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrimaryTaskSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServiceMixinProps",
		reflect.TypeOf((*CfnServiceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin",
		reflect.TypeOf((*CfnServicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.AdvancedConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AdvancedConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.AwsVpcConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AwsVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.CanaryConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CanaryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.CapacityProviderStrategyItemProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CapacityProviderStrategyItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.DeploymentAlarmsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DeploymentAlarmsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.DeploymentCircuitBreakerProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DeploymentCircuitBreakerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.DeploymentConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DeploymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.DeploymentControllerProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DeploymentControllerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.DeploymentLifecycleHookProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DeploymentLifecycleHookProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.EBSTagSpecificationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_EBSTagSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ForceNewDeploymentProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ForceNewDeploymentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.LinearConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_LinearConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.LoadBalancerProperty",
		reflect.TypeOf((*CfnServicePropsMixin_LoadBalancerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.PlacementConstraintProperty",
		reflect.TypeOf((*CfnServicePropsMixin_PlacementConstraintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.PlacementStrategyProperty",
		reflect.TypeOf((*CfnServicePropsMixin_PlacementStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.SecretProperty",
		reflect.TypeOf((*CfnServicePropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceConnectAccessLogConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectAccessLogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceConnectClientAliasProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectClientAliasProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceConnectConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceConnectServiceProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectServiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceConnectTestTrafficRulesHeaderProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectTestTrafficRulesHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceConnectTestTrafficRulesHeaderValueProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectTestTrafficRulesHeaderValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceConnectTestTrafficRulesProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectTestTrafficRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceConnectTlsCertificateAuthorityProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectTlsCertificateAuthorityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceConnectTlsConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceConnectTlsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceManagedEBSVolumeConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceManagedEBSVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceRegistryProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceRegistryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.ServiceVolumeConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.TimeoutConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_TimeoutConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin.VpcLatticeConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_VpcLatticeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionMixinProps",
		reflect.TypeOf((*CfnTaskDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTaskDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.AuthorizationConfigProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_AuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.ContainerDefinitionProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_ContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.ContainerDependencyProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_ContainerDependencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.DeviceProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_DeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.DockerVolumeConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_DockerVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.EFSVolumeConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_EFSVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.EnvironmentFileProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_EnvironmentFileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.EphemeralStorageProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_EphemeralStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.FSxAuthorizationConfigProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_FSxAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.FSxWindowsFileServerVolumeConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_FSxWindowsFileServerVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.FirelensConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_FirelensConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.HealthCheckProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_HealthCheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.HostEntryProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_HostEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.HostVolumePropertiesProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_HostVolumePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.InferenceAcceleratorProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_InferenceAcceleratorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.KernelCapabilitiesProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_KernelCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.KeyValuePairProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.LinuxParametersProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_LinuxParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.MountPointProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_MountPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.PortMappingProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_PortMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.ProxyConfigurationProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_ProxyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.RepositoryCredentialsProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_RepositoryCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.ResourceRequirementProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_ResourceRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.RestartPolicyProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_RestartPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.RuntimePlatformProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_RuntimePlatformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.SecretProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.SystemControlProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_SystemControlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.TaskDefinitionPlacementConstraintProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_TaskDefinitionPlacementConstraintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.TmpfsProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_TmpfsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.UlimitProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_UlimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.VolumeFromProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_VolumeFromProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin.VolumeProperty",
		reflect.TypeOf((*CfnTaskDefinitionPropsMixin_VolumeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetMixinProps",
		reflect.TypeOf((*CfnTaskSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin",
		reflect.TypeOf((*CfnTaskSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTaskSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin.AwsVpcConfigurationProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_AwsVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin.CapacityProviderStrategyItemProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_CapacityProviderStrategyItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin.LoadBalancerProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_LoadBalancerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin.ScaleProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_ScaleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin.ServiceRegistryProperty",
		reflect.TypeOf((*CfnTaskSetPropsMixin_ServiceRegistryProperty)(nil)).Elem(),
	)
}
