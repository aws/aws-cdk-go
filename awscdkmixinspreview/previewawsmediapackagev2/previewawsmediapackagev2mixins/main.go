package previewawsmediapackagev2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogs",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnChannelGroupEgressAccessLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogs",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnChannelGroupIngressAccessLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupLogsMixin",
		reflect.TypeOf((*CfnChannelGroupLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelGroupLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupMixinProps",
		reflect.TypeOf((*CfnChannelGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupPropsMixin",
		reflect.TypeOf((*CfnChannelGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelMixinProps",
		reflect.TypeOf((*CfnChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPolicyMixinProps",
		reflect.TypeOf((*CfnChannelPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPolicyPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPropsMixin.IngestEndpointProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_IngestEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPropsMixin.InputSwitchConfigurationProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_InputSwitchConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPropsMixin.OutputHeaderConfigurationProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_OutputHeaderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointMixinProps",
		reflect.TypeOf((*CfnOriginEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPolicyMixinProps",
		reflect.TypeOf((*CfnOriginEndpointPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPolicyPropsMixin",
		reflect.TypeOf((*CfnOriginEndpointPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginEndpointPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPolicyPropsMixin.CdnAuthConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPolicyPropsMixin_CdnAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashBaseUrlProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashBaseUrlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashDvbFontDownloadProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashDvbFontDownloadProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashDvbMetricsReportingProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashDvbMetricsReportingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashDvbSettingsProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashDvbSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashManifestConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashManifestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashProgramInformationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashProgramInformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashSubtitleConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashSubtitleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashTtmlConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashTtmlConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashUtcTimingProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashUtcTimingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.EncryptionContractConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_EncryptionContractConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.EncryptionMethodProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_EncryptionMethodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.EncryptionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_EncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.FilterConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_FilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.ForceEndpointErrorConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_ForceEndpointErrorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.HlsManifestConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_HlsManifestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.LowLatencyHlsManifestConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_LowLatencyHlsManifestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.MssManifestConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_MssManifestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.ScteDashProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_ScteDashProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.ScteHlsProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_ScteHlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.ScteProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_ScteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.SegmentProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_SegmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.SpekeKeyProviderProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_SpekeKeyProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.StartTagProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_StartTagProperty)(nil)).Elem(),
	)
}
