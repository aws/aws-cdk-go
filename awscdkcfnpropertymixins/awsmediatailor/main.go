package awsmediatailor

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnChannelMixinProps",
		reflect.TypeOf((*CfnChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnChannelPolicyMixinProps",
		reflect.TypeOf((*CfnChannelPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnChannelPolicyPropsMixin",
		reflect.TypeOf((*CfnChannelPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnChannelPropsMixin",
		reflect.TypeOf((*CfnChannelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnChannelPropsMixin.DashPlaylistSettingsProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_DashPlaylistSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnChannelPropsMixin.HlsPlaylistSettingsProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_HlsPlaylistSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnChannelPropsMixin.LogConfigurationForChannelProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_LogConfigurationForChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnChannelPropsMixin.RequestOutputItemProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_RequestOutputItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnChannelPropsMixin.SlateSourceProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_SlateSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnChannelPropsMixin.TimeShiftConfigurationProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_TimeShiftConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnLiveSourceMixinProps",
		reflect.TypeOf((*CfnLiveSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnLiveSourcePropsMixin",
		reflect.TypeOf((*CfnLiveSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLiveSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnLiveSourcePropsMixin.HttpPackageConfigurationProperty",
		reflect.TypeOf((*CfnLiveSourcePropsMixin_HttpPackageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationMixinProps",
		reflect.TypeOf((*CfnPlaybackConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPlaybackConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.AdConditioningConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdConditioningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.AdDecisionServerConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdDecisionServerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.AdMarkerPassthroughProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdMarkerPassthroughProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.AdsInteractionLogProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdsInteractionLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.AvailSuppressionProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AvailSuppressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.BumperProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_BumperProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.CdnConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_CdnConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.DashConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_DashConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.HlsConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_HlsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.HttpRequestProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_HttpRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.LivePreRollConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_LivePreRollConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.ManifestProcessingRulesProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_ManifestProcessingRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPlaybackConfigurationPropsMixin.ManifestServiceInteractionLogProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_ManifestServiceInteractionLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnSourceLocationMixinProps",
		reflect.TypeOf((*CfnSourceLocationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnSourceLocationPropsMixin",
		reflect.TypeOf((*CfnSourceLocationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSourceLocationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnSourceLocationPropsMixin.AccessConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_AccessConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnSourceLocationPropsMixin.DefaultSegmentDeliveryConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_DefaultSegmentDeliveryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnSourceLocationPropsMixin.HttpConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_HttpConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnSourceLocationPropsMixin.SecretsManagerAccessTokenConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_SecretsManagerAccessTokenConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnSourceLocationPropsMixin.SegmentDeliveryConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_SegmentDeliveryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnVodSourceMixinProps",
		reflect.TypeOf((*CfnVodSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnVodSourcePropsMixin",
		reflect.TypeOf((*CfnVodSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVodSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnVodSourcePropsMixin.HttpPackageConfigurationProperty",
		reflect.TypeOf((*CfnVodSourcePropsMixin_HttpPackageConfigurationProperty)(nil)).Elem(),
	)
}
