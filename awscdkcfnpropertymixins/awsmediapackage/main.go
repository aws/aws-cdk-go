package awsmediapackage

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnAssetMixinProps",
		reflect.TypeOf((*CfnAssetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnAssetPropsMixin",
		reflect.TypeOf((*CfnAssetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnAssetPropsMixin.EgressEndpointProperty",
		reflect.TypeOf((*CfnAssetPropsMixin_EgressEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnChannelMixinProps",
		reflect.TypeOf((*CfnChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnChannelPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnChannelPropsMixin.HlsIngestProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_HlsIngestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnChannelPropsMixin.IngestEndpointProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_IngestEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnChannelPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointMixinProps",
		reflect.TypeOf((*CfnOriginEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.AuthorizationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_AuthorizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.CmafEncryptionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_CmafEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.CmafPackageProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_CmafPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.DashEncryptionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.DashPackageProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.EncryptionContractConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_EncryptionContractConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.HlsEncryptionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_HlsEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.HlsManifestProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_HlsManifestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.HlsPackageProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_HlsPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.MssEncryptionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_MssEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.MssPackageProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_MssPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.SpekeKeyProviderProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_SpekeKeyProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnOriginEndpointPropsMixin.StreamSelectionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_StreamSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationMixinProps",
		reflect.TypeOf((*CfnPackagingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPackagingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.CmafEncryptionProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_CmafEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.CmafPackageProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_CmafPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.DashEncryptionProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_DashEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.DashManifestProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_DashManifestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.DashPackageProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_DashPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.EncryptionContractConfigurationProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_EncryptionContractConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.HlsEncryptionProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_HlsEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.HlsManifestProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_HlsManifestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.HlsPackageProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_HlsPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.MssEncryptionProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_MssEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.MssManifestProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_MssManifestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.MssPackageProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_MssPackageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.SpekeKeyProviderProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_SpekeKeyProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingConfigurationPropsMixin.StreamSelectionProperty",
		reflect.TypeOf((*CfnPackagingConfigurationPropsMixin_StreamSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingGroupMixinProps",
		reflect.TypeOf((*CfnPackagingGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingGroupPropsMixin",
		reflect.TypeOf((*CfnPackagingGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPackagingGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingGroupPropsMixin.AuthorizationProperty",
		reflect.TypeOf((*CfnPackagingGroupPropsMixin_AuthorizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnPackagingGroupPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnPackagingGroupPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
}
