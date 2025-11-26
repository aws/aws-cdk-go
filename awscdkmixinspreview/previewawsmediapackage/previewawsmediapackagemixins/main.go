package previewawsmediapackagemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnAssetMixinProps",
		reflect.TypeOf((*CfnAssetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnAssetPropsMixin",
		reflect.TypeOf((*CfnAssetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnAssetPropsMixin.EgressEndpointProperty",
		reflect.TypeOf((*CfnAssetPropsMixin_EgressEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnChannelMixinProps",
		reflect.TypeOf((*CfnChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnChannelPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnChannelPropsMixin.HlsIngestProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_HlsIngestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnChannelPropsMixin.IngestEndpointProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_IngestEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnChannelPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointMixinProps",
		reflect.TypeOf((*CfnOriginEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.AuthorizationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_AuthorizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.CmafEncryptionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_CmafEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.CmafPackageProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_CmafPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.DashEncryptionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.DashPackageProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.EncryptionContractConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_EncryptionContractConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.HlsEncryptionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_HlsEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.HlsManifestProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_HlsManifestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.HlsPackageProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_HlsPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.MssEncryptionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_MssEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.MssPackageProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_MssPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.SpekeKeyProviderProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_SpekeKeyProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnOriginEndpointPropsMixin.StreamSelectionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_StreamSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationMixinProps",
		reflect.TypeOf((*CfnPackagingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPackagingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.CmafEncryptionProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_CmafEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.CmafPackageProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_CmafPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.DashEncryptionProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_DashEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.DashManifestProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_DashManifestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.DashPackageProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_DashPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.EncryptionContractConfigurationProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_EncryptionContractConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.HlsEncryptionProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_HlsEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.HlsManifestProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_HlsManifestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.HlsPackageProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_HlsPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.MssEncryptionProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_MssEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.MssManifestProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_MssManifestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.MssPackageProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_MssPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.SpekeKeyProviderProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_SpekeKeyProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingConfigurationPropsMixin.StreamSelectionProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_StreamSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingGroupMixinProps",
		reflect.TypeOf((*CfnPackagingGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingGroupPropsMixin",
		reflect.TypeOf((*CfnPackagingGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPackagingGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingGroupPropsMixin.AuthorizationProperty",
		reflect.TypeOf((*CfnPackagingGroupPropsMixin_AuthorizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.mixins.CfnPackagingGroupPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnPackagingGroupPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
}
