package awsgamelift

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnAliasMixinProps",
		reflect.TypeOf((*CfnAliasMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnAliasPropsMixin",
		reflect.TypeOf((*CfnAliasPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAliasPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnAliasPropsMixin.RoutingStrategyProperty",
		reflect.TypeOf((*CfnAliasPropsMixin_RoutingStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnBuildMixinProps",
		reflect.TypeOf((*CfnBuildMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnBuildPropsMixin",
		reflect.TypeOf((*CfnBuildPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBuildPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnBuildPropsMixin.StorageLocationProperty",
		reflect.TypeOf((*CfnBuildPropsMixin_StorageLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetMixinProps",
		reflect.TypeOf((*CfnContainerFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin",
		reflect.TypeOf((*CfnContainerFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContainerFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.ConnectionPortRangeProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_ConnectionPortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.DeploymentConfigurationProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_DeploymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.DeploymentDetailsProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_DeploymentDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.GameSessionCreationLimitPolicyProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_GameSessionCreationLimitPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.IpPermissionProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_IpPermissionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.LocationCapacityProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_LocationCapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.LocationConfigurationProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_LocationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.ManagedCapacityConfigurationProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_ManagedCapacityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.ScalingPolicyProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_ScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerFleetPropsMixin.TargetConfigurationProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_TargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionMixinProps",
		reflect.TypeOf((*CfnContainerGroupDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionPropsMixin",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContainerGroupDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionPropsMixin.ContainerDependencyProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_ContainerDependencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionPropsMixin.ContainerEnvironmentProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_ContainerEnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionPropsMixin.ContainerHealthCheckProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_ContainerHealthCheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionPropsMixin.ContainerMountPointProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_ContainerMountPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionPropsMixin.ContainerPortRangeProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_ContainerPortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionPropsMixin.GameServerContainerDefinitionProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_GameServerContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionPropsMixin.LinuxCapabilitiesProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_LinuxCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionPropsMixin.PortConfigurationProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_PortConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnContainerGroupDefinitionPropsMixin.SupportContainerDefinitionProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_SupportContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetMixinProps",
		reflect.TypeOf((*CfnFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin",
		reflect.TypeOf((*CfnFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.AnywhereConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_AnywhereConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.CertificateConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_CertificateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.IpPermissionProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_IpPermissionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.LocationCapacityProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_LocationCapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.LocationConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_LocationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.ManagedCapacityConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ManagedCapacityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.PlayerGatewayConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_PlayerGatewayConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.ResourceCreationLimitPolicyProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ResourceCreationLimitPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.RuntimeConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_RuntimeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.ScalingPolicyProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.ServerProcessProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ServerProcessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnFleetPropsMixin.TargetConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_TargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameServerGroupMixinProps",
		reflect.TypeOf((*CfnGameServerGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameServerGroupPropsMixin",
		reflect.TypeOf((*CfnGameServerGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGameServerGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameServerGroupPropsMixin.AutoScalingPolicyProperty",
		reflect.TypeOf((*CfnGameServerGroupPropsMixin_AutoScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameServerGroupPropsMixin.InstanceDefinitionProperty",
		reflect.TypeOf((*CfnGameServerGroupPropsMixin_InstanceDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameServerGroupPropsMixin.LaunchTemplateProperty",
		reflect.TypeOf((*CfnGameServerGroupPropsMixin_LaunchTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameServerGroupPropsMixin.TargetTrackingConfigurationProperty",
		reflect.TypeOf((*CfnGameServerGroupPropsMixin_TargetTrackingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameSessionQueueMixinProps",
		reflect.TypeOf((*CfnGameSessionQueueMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameSessionQueuePropsMixin",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGameSessionQueuePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameSessionQueuePropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameSessionQueuePropsMixin.FilterConfigurationProperty",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin_FilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameSessionQueuePropsMixin.GameSessionQueueDestinationProperty",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin_GameSessionQueueDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameSessionQueuePropsMixin.PlayerLatencyPolicyProperty",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin_PlayerLatencyPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnGameSessionQueuePropsMixin.PriorityConfigurationProperty",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin_PriorityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnLocationMixinProps",
		reflect.TypeOf((*CfnLocationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnLocationPropsMixin",
		reflect.TypeOf((*CfnLocationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnMatchmakingConfigurationMixinProps",
		reflect.TypeOf((*CfnMatchmakingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnMatchmakingConfigurationPropsMixin",
		reflect.TypeOf((*CfnMatchmakingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMatchmakingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnMatchmakingConfigurationPropsMixin.GamePropertyProperty",
		reflect.TypeOf((*CfnMatchmakingConfigurationPropsMixin_GamePropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnMatchmakingRuleSetMixinProps",
		reflect.TypeOf((*CfnMatchmakingRuleSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnMatchmakingRuleSetPropsMixin",
		reflect.TypeOf((*CfnMatchmakingRuleSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMatchmakingRuleSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnScriptMixinProps",
		reflect.TypeOf((*CfnScriptMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnScriptPropsMixin",
		reflect.TypeOf((*CfnScriptPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScriptPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_gamelift.CfnScriptPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnScriptPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
}
