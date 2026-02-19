package previewawseksmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAccessEntryMixinProps",
		reflect.TypeOf((*CfnAccessEntryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAccessEntryPropsMixin",
		reflect.TypeOf((*CfnAccessEntryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessEntryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAccessEntryPropsMixin.AccessPolicyProperty",
		reflect.TypeOf((*CfnAccessEntryPropsMixin_AccessPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAccessEntryPropsMixin.AccessScopeProperty",
		reflect.TypeOf((*CfnAccessEntryPropsMixin_AccessScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAddonMixinProps",
		reflect.TypeOf((*CfnAddonMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAddonPropsMixin",
		reflect.TypeOf((*CfnAddonPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAddonPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAddonPropsMixin.NamespaceConfigProperty",
		reflect.TypeOf((*CfnAddonPropsMixin_NamespaceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnAddonPropsMixin.PodIdentityAssociationProperty",
		reflect.TypeOf((*CfnAddonPropsMixin_PodIdentityAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityMixinProps",
		reflect.TypeOf((*CfnCapabilityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin",
		reflect.TypeOf((*CfnCapabilityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapabilityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin.ArgoCdProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_ArgoCdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin.ArgoCdRoleMappingProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_ArgoCdRoleMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin.AwsIdcProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_AwsIdcProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin.CapabilityConfigurationProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_CapabilityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin.NetworkAccessProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_NetworkAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityPropsMixin.SsoIdentityProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_SsoIdentityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogs",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeBlockStorageLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogs",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeComputeLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogs",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeIpamLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogs",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeLoadBalancingLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterLogsMixin",
		reflect.TypeOf((*CfnClusterLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.AccessConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AccessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.BlockStorageProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_BlockStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.ClusterLoggingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterLoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.ComputeConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ComputeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.ControlPlanePlacementProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ControlPlanePlacementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.ControlPlaneScalingConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ControlPlaneScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.ElasticLoadBalancingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ElasticLoadBalancingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.EncryptionConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EncryptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.KubernetesNetworkConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_KubernetesNetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.LoggingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.LoggingTypeConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_LoggingTypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.OutpostConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OutpostConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.ProviderProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.RemoteNetworkConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_RemoteNetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.RemoteNodeNetworkProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_RemoteNodeNetworkProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.RemotePodNetworkProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_RemotePodNetworkProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.ResourcesVpcConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ResourcesVpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.StorageConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_StorageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.UpgradePolicyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_UpgradePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin.ZonalShiftConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ZonalShiftConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnFargateProfileMixinProps",
		reflect.TypeOf((*CfnFargateProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnFargateProfilePropsMixin",
		reflect.TypeOf((*CfnFargateProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFargateProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnFargateProfilePropsMixin.LabelProperty",
		reflect.TypeOf((*CfnFargateProfilePropsMixin_LabelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnFargateProfilePropsMixin.SelectorProperty",
		reflect.TypeOf((*CfnFargateProfilePropsMixin_SelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnIdentityProviderConfigMixinProps",
		reflect.TypeOf((*CfnIdentityProviderConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnIdentityProviderConfigPropsMixin",
		reflect.TypeOf((*CfnIdentityProviderConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentityProviderConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnIdentityProviderConfigPropsMixin.OidcIdentityProviderConfigProperty",
		reflect.TypeOf((*CfnIdentityProviderConfigPropsMixin_OidcIdentityProviderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnIdentityProviderConfigPropsMixin.RequiredClaimProperty",
		reflect.TypeOf((*CfnIdentityProviderConfigPropsMixin_RequiredClaimProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupMixinProps",
		reflect.TypeOf((*CfnNodegroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin",
		reflect.TypeOf((*CfnNodegroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNodegroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin.LaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_LaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin.NodeRepairConfigOverridesProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_NodeRepairConfigOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin.NodeRepairConfigProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_NodeRepairConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin.RemoteAccessProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_RemoteAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin.ScalingConfigProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_ScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin.TaintProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_TaintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnNodegroupPropsMixin.UpdateConfigProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_UpdateConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnPodIdentityAssociationMixinProps",
		reflect.TypeOf((*CfnPodIdentityAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnPodIdentityAssociationPropsMixin",
		reflect.TypeOf((*CfnPodIdentityAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPodIdentityAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
