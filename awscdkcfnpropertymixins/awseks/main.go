package awseks

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnAccessEntryMixinProps",
		reflect.TypeOf((*CfnAccessEntryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnAccessEntryPropsMixin",
		reflect.TypeOf((*CfnAccessEntryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessEntryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnAccessEntryPropsMixin.AccessPolicyProperty",
		reflect.TypeOf((*CfnAccessEntryPropsMixin_AccessPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnAccessEntryPropsMixin.AccessScopeProperty",
		reflect.TypeOf((*CfnAccessEntryPropsMixin_AccessScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnAddonMixinProps",
		reflect.TypeOf((*CfnAddonMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnAddonPropsMixin",
		reflect.TypeOf((*CfnAddonPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAddonPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnAddonPropsMixin.NamespaceConfigProperty",
		reflect.TypeOf((*CfnAddonPropsMixin_NamespaceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnAddonPropsMixin.PodIdentityAssociationProperty",
		reflect.TypeOf((*CfnAddonPropsMixin_PodIdentityAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnCapabilityMixinProps",
		reflect.TypeOf((*CfnCapabilityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnCapabilityPropsMixin",
		reflect.TypeOf((*CfnCapabilityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapabilityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnCapabilityPropsMixin.ArgoCdProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_ArgoCdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnCapabilityPropsMixin.ArgoCdRoleMappingProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_ArgoCdRoleMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnCapabilityPropsMixin.AwsIdcProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_AwsIdcProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnCapabilityPropsMixin.CapabilityConfigurationProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_CapabilityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnCapabilityPropsMixin.NetworkAccessProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_NetworkAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnCapabilityPropsMixin.SsoIdentityProperty",
		reflect.TypeOf((*CfnCapabilityPropsMixin_SsoIdentityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.AccessConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AccessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.BlockStorageProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_BlockStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.ClusterLoggingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClusterLoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.ComputeConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ComputeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.ControlPlanePlacementProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ControlPlanePlacementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.ControlPlaneScalingConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ControlPlaneScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.ElasticLoadBalancingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ElasticLoadBalancingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.EncryptionConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EncryptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.EtcdPlacementProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EtcdPlacementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.KubernetesNetworkConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_KubernetesNetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.LoggingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.LoggingTypeConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_LoggingTypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.OutpostConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OutpostConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.ProviderProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.RemoteNetworkConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_RemoteNetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.RemoteNodeNetworkProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_RemoteNodeNetworkProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.RemotePodNetworkProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_RemotePodNetworkProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.ResourcesVpcConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ResourcesVpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.StorageConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_StorageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.UpgradePolicyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_UpgradePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnClusterPropsMixin.ZonalShiftConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ZonalShiftConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnFargateProfileMixinProps",
		reflect.TypeOf((*CfnFargateProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnFargateProfilePropsMixin",
		reflect.TypeOf((*CfnFargateProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFargateProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnFargateProfilePropsMixin.LabelProperty",
		reflect.TypeOf((*CfnFargateProfilePropsMixin_LabelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnFargateProfilePropsMixin.SelectorProperty",
		reflect.TypeOf((*CfnFargateProfilePropsMixin_SelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnIdentityProviderConfigMixinProps",
		reflect.TypeOf((*CfnIdentityProviderConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnIdentityProviderConfigPropsMixin",
		reflect.TypeOf((*CfnIdentityProviderConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentityProviderConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnIdentityProviderConfigPropsMixin.OidcIdentityProviderConfigProperty",
		reflect.TypeOf((*CfnIdentityProviderConfigPropsMixin_OidcIdentityProviderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnIdentityProviderConfigPropsMixin.RequiredClaimProperty",
		reflect.TypeOf((*CfnIdentityProviderConfigPropsMixin_RequiredClaimProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnNodegroupMixinProps",
		reflect.TypeOf((*CfnNodegroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnNodegroupPropsMixin",
		reflect.TypeOf((*CfnNodegroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNodegroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnNodegroupPropsMixin.LaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_LaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnNodegroupPropsMixin.NodeRepairConfigOverridesProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_NodeRepairConfigOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnNodegroupPropsMixin.NodeRepairConfigProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_NodeRepairConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnNodegroupPropsMixin.RemoteAccessProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_RemoteAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnNodegroupPropsMixin.ScalingConfigProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_ScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnNodegroupPropsMixin.TaintProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_TaintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnNodegroupPropsMixin.UpdateConfigProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_UpdateConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnNodegroupPropsMixin.WarmPoolConfigProperty",
		reflect.TypeOf((*CfnNodegroupPropsMixin_WarmPoolConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnPodIdentityAssociationMixinProps",
		reflect.TypeOf((*CfnPodIdentityAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_eks.CfnPodIdentityAssociationPropsMixin",
		reflect.TypeOf((*CfnPodIdentityAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPodIdentityAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
