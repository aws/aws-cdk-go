package mixinsawsgamelift

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnAliasMixinProps",
		reflect.TypeOf((*CfnAliasMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnAliasPropsMixin",
		reflect.TypeOf((*CfnAliasPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAliasPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnAliasPropsMixin.RoutingStrategyProperty",
		reflect.TypeOf((*CfnAliasPropsMixin_RoutingStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnBuildMixinProps",
		reflect.TypeOf((*CfnBuildMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnBuildPropsMixin",
		reflect.TypeOf((*CfnBuildPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBuildPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnBuildPropsMixin.StorageLocationProperty",
		reflect.TypeOf((*CfnBuildPropsMixin_StorageLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetMixinProps",
		reflect.TypeOf((*CfnContainerFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin",
		reflect.TypeOf((*CfnContainerFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContainerFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin.ConnectionPortRangeProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_ConnectionPortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin.DeploymentConfigurationProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_DeploymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin.DeploymentDetailsProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_DeploymentDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin.GameSessionCreationLimitPolicyProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_GameSessionCreationLimitPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin.IpPermissionProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_IpPermissionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin.LocationCapacityProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_LocationCapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin.LocationConfigurationProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_LocationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin.ScalingPolicyProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_ScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin.TargetConfigurationProperty",
		reflect.TypeOf((*CfnContainerFleetPropsMixin_TargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionMixinProps",
		reflect.TypeOf((*CfnContainerGroupDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContainerGroupDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin.ContainerDependencyProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_ContainerDependencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin.ContainerEnvironmentProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_ContainerEnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin.ContainerHealthCheckProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_ContainerHealthCheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin.ContainerMountPointProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_ContainerMountPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin.ContainerPortRangeProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_ContainerPortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin.GameServerContainerDefinitionProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_GameServerContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin.PortConfigurationProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_PortConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin.SupportContainerDefinitionProperty",
		reflect.TypeOf((*CfnContainerGroupDefinitionPropsMixin_SupportContainerDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetMixinProps",
		reflect.TypeOf((*CfnFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin",
		reflect.TypeOf((*CfnFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin.AnywhereConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_AnywhereConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin.CertificateConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_CertificateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin.IpPermissionProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_IpPermissionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin.LocationCapacityProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_LocationCapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin.LocationConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_LocationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin.ResourceCreationLimitPolicyProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ResourceCreationLimitPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin.RuntimeConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_RuntimeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin.ScalingPolicyProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin.ServerProcessProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ServerProcessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin.TargetConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_TargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameServerGroupMixinProps",
		reflect.TypeOf((*CfnGameServerGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameServerGroupPropsMixin",
		reflect.TypeOf((*CfnGameServerGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGameServerGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameServerGroupPropsMixin.AutoScalingPolicyProperty",
		reflect.TypeOf((*CfnGameServerGroupPropsMixin_AutoScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameServerGroupPropsMixin.InstanceDefinitionProperty",
		reflect.TypeOf((*CfnGameServerGroupPropsMixin_InstanceDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameServerGroupPropsMixin.LaunchTemplateProperty",
		reflect.TypeOf((*CfnGameServerGroupPropsMixin_LaunchTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameServerGroupPropsMixin.TargetTrackingConfigurationProperty",
		reflect.TypeOf((*CfnGameServerGroupPropsMixin_TargetTrackingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueMixinProps",
		reflect.TypeOf((*CfnGameSessionQueueMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueuePropsMixin",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGameSessionQueuePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueuePropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueuePropsMixin.FilterConfigurationProperty",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin_FilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueuePropsMixin.GameSessionQueueDestinationProperty",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin_GameSessionQueueDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueuePropsMixin.PlayerLatencyPolicyProperty",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin_PlayerLatencyPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueuePropsMixin.PriorityConfigurationProperty",
		reflect.TypeOf((*CfnGameSessionQueuePropsMixin_PriorityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnLocationMixinProps",
		reflect.TypeOf((*CfnLocationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnLocationPropsMixin",
		reflect.TypeOf((*CfnLocationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnMatchmakingConfigurationMixinProps",
		reflect.TypeOf((*CfnMatchmakingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnMatchmakingConfigurationPropsMixin",
		reflect.TypeOf((*CfnMatchmakingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMatchmakingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnMatchmakingConfigurationPropsMixin.GamePropertyProperty",
		reflect.TypeOf((*CfnMatchmakingConfigurationPropsMixin_GamePropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnMatchmakingRuleSetMixinProps",
		reflect.TypeOf((*CfnMatchmakingRuleSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnMatchmakingRuleSetPropsMixin",
		reflect.TypeOf((*CfnMatchmakingRuleSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMatchmakingRuleSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnScriptMixinProps",
		reflect.TypeOf((*CfnScriptMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnScriptPropsMixin",
		reflect.TypeOf((*CfnScriptPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScriptPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnScriptPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnScriptPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
}
