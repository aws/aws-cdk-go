package previewawsmediatailormixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelMixinProps",
		reflect.TypeOf((*CfnChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPolicyMixinProps",
		reflect.TypeOf((*CfnChannelPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPolicyPropsMixin",
		reflect.TypeOf((*CfnChannelPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin",
		reflect.TypeOf((*CfnChannelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.DashPlaylistSettingsProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_DashPlaylistSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.HlsPlaylistSettingsProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_HlsPlaylistSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.LogConfigurationForChannelProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_LogConfigurationForChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.RequestOutputItemProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_RequestOutputItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.SlateSourceProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_SlateSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.TimeShiftConfigurationProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_TimeShiftConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnLiveSourceMixinProps",
		reflect.TypeOf((*CfnLiveSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnLiveSourcePropsMixin",
		reflect.TypeOf((*CfnLiveSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLiveSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnLiveSourcePropsMixin.HttpPackageConfigurationProperty",
		reflect.TypeOf((*CfnLiveSourcePropsMixin_HttpPackageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogs",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnPlaybackConfigurationAdDecisionServerLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationLogsMixin",
		reflect.TypeOf((*CfnPlaybackConfigurationLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPlaybackConfigurationLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogs",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationMixinProps",
		reflect.TypeOf((*CfnPlaybackConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPlaybackConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.AdConditioningConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdConditioningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.AdDecisionServerConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdDecisionServerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.AdMarkerPassthroughProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdMarkerPassthroughProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.AdsInteractionLogProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdsInteractionLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.AvailSuppressionProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AvailSuppressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.BumperProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_BumperProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.CdnConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_CdnConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.DashConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_DashConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.HlsConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_HlsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.HttpRequestProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_HttpRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.LivePreRollConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_LivePreRollConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.ManifestProcessingRulesProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_ManifestProcessingRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.ManifestServiceInteractionLogProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_ManifestServiceInteractionLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogs",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnPlaybackConfigurationTranscodeLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationMixinProps",
		reflect.TypeOf((*CfnSourceLocationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin",
		reflect.TypeOf((*CfnSourceLocationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSourceLocationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin.AccessConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_AccessConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin.DefaultSegmentDeliveryConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_DefaultSegmentDeliveryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin.HttpConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_HttpConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin.SecretsManagerAccessTokenConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_SecretsManagerAccessTokenConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin.SegmentDeliveryConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_SegmentDeliveryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnVodSourceMixinProps",
		reflect.TypeOf((*CfnVodSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnVodSourcePropsMixin",
		reflect.TypeOf((*CfnVodSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVodSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnVodSourcePropsMixin.HttpPackageConfigurationProperty",
		reflect.TypeOf((*CfnVodSourcePropsMixin_HttpPackageConfigurationProperty)(nil)).Elem(),
	)
}
