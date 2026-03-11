package awsqbusiness

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnApplicationPropsMixin.AttachmentsConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_AttachmentsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnApplicationPropsMixin.AutoSubscriptionConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_AutoSubscriptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnApplicationPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnApplicationPropsMixin.PersonalizationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_PersonalizationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnApplicationPropsMixin.QAppsConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_QAppsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnApplicationPropsMixin.QuickSightConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_QuickSightConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataAccessorMixinProps",
		reflect.TypeOf((*CfnDataAccessorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataAccessorPropsMixin",
		reflect.TypeOf((*CfnDataAccessorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataAccessorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataAccessorPropsMixin.ActionConfigurationProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_ActionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataAccessorPropsMixin.ActionFilterConfigurationProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_ActionFilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataAccessorPropsMixin.AttributeFilterProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_AttributeFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataAccessorPropsMixin.DataAccessorAuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_DataAccessorAuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataAccessorPropsMixin.DataAccessorAuthenticationDetailProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_DataAccessorAuthenticationDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataAccessorPropsMixin.DataAccessorIdcTrustedTokenIssuerConfigurationProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_DataAccessorIdcTrustedTokenIssuerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataAccessorPropsMixin.DocumentAttributeProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_DocumentAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataAccessorPropsMixin.DocumentAttributeValueProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_DocumentAttributeValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourceMixinProps",
		reflect.TypeOf((*CfnDataSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin",
		reflect.TypeOf((*CfnDataSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.AudioExtractionConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_AudioExtractionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.DataSourceVpcConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DataSourceVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.DocumentAttributeConditionProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DocumentAttributeConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.DocumentAttributeTargetProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DocumentAttributeTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.DocumentAttributeValueProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DocumentAttributeValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.DocumentEnrichmentConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DocumentEnrichmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.HookConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_HookConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.ImageExtractionConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_ImageExtractionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.InlineDocumentEnrichmentConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_InlineDocumentEnrichmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.MediaExtractionConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_MediaExtractionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnDataSourcePropsMixin.VideoExtractionConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_VideoExtractionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnIndexMixinProps",
		reflect.TypeOf((*CfnIndexMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnIndexPropsMixin",
		reflect.TypeOf((*CfnIndexPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIndexPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnIndexPropsMixin.DocumentAttributeConfigurationProperty",
		reflect.TypeOf((*CfnIndexPropsMixin_DocumentAttributeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnIndexPropsMixin.IndexCapacityConfigurationProperty",
		reflect.TypeOf((*CfnIndexPropsMixin_IndexCapacityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnIndexPropsMixin.IndexStatisticsProperty",
		reflect.TypeOf((*CfnIndexPropsMixin_IndexStatisticsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnIndexPropsMixin.TextDocumentStatisticsProperty",
		reflect.TypeOf((*CfnIndexPropsMixin_TextDocumentStatisticsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPermissionMixinProps",
		reflect.TypeOf((*CfnPermissionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPermissionPropsMixin",
		reflect.TypeOf((*CfnPermissionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPermissionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPermissionPropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnPermissionPropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPluginMixinProps",
		reflect.TypeOf((*CfnPluginMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPluginPropsMixin",
		reflect.TypeOf((*CfnPluginPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPluginPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPluginPropsMixin.APISchemaProperty",
		reflect.TypeOf((*CfnPluginPropsMixin_APISchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPluginPropsMixin.BasicAuthConfigurationProperty",
		reflect.TypeOf((*CfnPluginPropsMixin_BasicAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPluginPropsMixin.CustomPluginConfigurationProperty",
		reflect.TypeOf((*CfnPluginPropsMixin_CustomPluginConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPluginPropsMixin.OAuth2ClientCredentialConfigurationProperty",
		reflect.TypeOf((*CfnPluginPropsMixin_OAuth2ClientCredentialConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPluginPropsMixin.PluginAuthConfigurationProperty",
		reflect.TypeOf((*CfnPluginPropsMixin_PluginAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnPluginPropsMixin.S3Property",
		reflect.TypeOf((*CfnPluginPropsMixin_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnRetrieverMixinProps",
		reflect.TypeOf((*CfnRetrieverMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnRetrieverPropsMixin",
		reflect.TypeOf((*CfnRetrieverPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRetrieverPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnRetrieverPropsMixin.KendraIndexConfigurationProperty",
		reflect.TypeOf((*CfnRetrieverPropsMixin_KendraIndexConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnRetrieverPropsMixin.NativeIndexConfigurationProperty",
		reflect.TypeOf((*CfnRetrieverPropsMixin_NativeIndexConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnRetrieverPropsMixin.RetrieverConfigurationProperty",
		reflect.TypeOf((*CfnRetrieverPropsMixin_RetrieverConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnWebExperienceMixinProps",
		reflect.TypeOf((*CfnWebExperienceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnWebExperiencePropsMixin",
		reflect.TypeOf((*CfnWebExperiencePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWebExperiencePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnWebExperiencePropsMixin.BrowserExtensionConfigurationProperty",
		reflect.TypeOf((*CfnWebExperiencePropsMixin_BrowserExtensionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnWebExperiencePropsMixin.CustomizationConfigurationProperty",
		reflect.TypeOf((*CfnWebExperiencePropsMixin_CustomizationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnWebExperiencePropsMixin.IdentityProviderConfigurationProperty",
		reflect.TypeOf((*CfnWebExperiencePropsMixin_IdentityProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnWebExperiencePropsMixin.OpenIDConnectProviderConfigurationProperty",
		reflect.TypeOf((*CfnWebExperiencePropsMixin_OpenIDConnectProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_qbusiness.CfnWebExperiencePropsMixin.SamlProviderConfigurationProperty",
		reflect.TypeOf((*CfnWebExperiencePropsMixin_SamlProviderConfigurationProperty)(nil)).Elem(),
	)
}
