package previewawsqbusinessmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogs",
		reflect.TypeOf((*CfnApplicationEventLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnApplicationEventLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationLogsMixin",
		reflect.TypeOf((*CfnApplicationLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin.AttachmentsConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_AttachmentsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin.AutoSubscriptionConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_AutoSubscriptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin.PersonalizationConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_PersonalizationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin.QAppsConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_QAppsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin.QuickSightConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_QuickSightConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogs",
		reflect.TypeOf((*CfnApplicationSyncJobLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnApplicationSyncJobLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorMixinProps",
		reflect.TypeOf((*CfnDataAccessorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin",
		reflect.TypeOf((*CfnDataAccessorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataAccessorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin.ActionConfigurationProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_ActionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin.ActionFilterConfigurationProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_ActionFilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin.AttributeFilterProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_AttributeFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin.DataAccessorAuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_DataAccessorAuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin.DataAccessorAuthenticationDetailProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_DataAccessorAuthenticationDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin.DataAccessorIdcTrustedTokenIssuerConfigurationProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_DataAccessorIdcTrustedTokenIssuerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin.DocumentAttributeProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_DocumentAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataAccessorPropsMixin.DocumentAttributeValueProperty",
		reflect.TypeOf((*CfnDataAccessorPropsMixin_DocumentAttributeValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourceMixinProps",
		reflect.TypeOf((*CfnDataSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin",
		reflect.TypeOf((*CfnDataSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.AudioExtractionConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_AudioExtractionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.DataSourceVpcConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DataSourceVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.DocumentAttributeConditionProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DocumentAttributeConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.DocumentAttributeTargetProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DocumentAttributeTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.DocumentAttributeValueProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DocumentAttributeValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.DocumentEnrichmentConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_DocumentEnrichmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.HookConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_HookConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.ImageExtractionConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_ImageExtractionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.InlineDocumentEnrichmentConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_InlineDocumentEnrichmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.MediaExtractionConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_MediaExtractionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin.VideoExtractionConfigurationProperty",
		reflect.TypeOf((*CfnDataSourcePropsMixin_VideoExtractionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnIndexMixinProps",
		reflect.TypeOf((*CfnIndexMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnIndexPropsMixin",
		reflect.TypeOf((*CfnIndexPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIndexPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnIndexPropsMixin.DocumentAttributeConfigurationProperty",
		reflect.TypeOf((*CfnIndexPropsMixin_DocumentAttributeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnIndexPropsMixin.IndexCapacityConfigurationProperty",
		reflect.TypeOf((*CfnIndexPropsMixin_IndexCapacityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnIndexPropsMixin.IndexStatisticsProperty",
		reflect.TypeOf((*CfnIndexPropsMixin_IndexStatisticsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnIndexPropsMixin.TextDocumentStatisticsProperty",
		reflect.TypeOf((*CfnIndexPropsMixin_TextDocumentStatisticsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPermissionMixinProps",
		reflect.TypeOf((*CfnPermissionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPermissionPropsMixin",
		reflect.TypeOf((*CfnPermissionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPermissionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPermissionPropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnPermissionPropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginMixinProps",
		reflect.TypeOf((*CfnPluginMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin",
		reflect.TypeOf((*CfnPluginPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPluginPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin.APISchemaProperty",
		reflect.TypeOf((*CfnPluginPropsMixin_APISchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin.BasicAuthConfigurationProperty",
		reflect.TypeOf((*CfnPluginPropsMixin_BasicAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin.CustomPluginConfigurationProperty",
		reflect.TypeOf((*CfnPluginPropsMixin_CustomPluginConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin.OAuth2ClientCredentialConfigurationProperty",
		reflect.TypeOf((*CfnPluginPropsMixin_OAuth2ClientCredentialConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin.PluginAuthConfigurationProperty",
		reflect.TypeOf((*CfnPluginPropsMixin_PluginAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPluginPropsMixin.S3Property",
		reflect.TypeOf((*CfnPluginPropsMixin_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnRetrieverMixinProps",
		reflect.TypeOf((*CfnRetrieverMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnRetrieverPropsMixin",
		reflect.TypeOf((*CfnRetrieverPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRetrieverPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnRetrieverPropsMixin.KendraIndexConfigurationProperty",
		reflect.TypeOf((*CfnRetrieverPropsMixin_KendraIndexConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnRetrieverPropsMixin.NativeIndexConfigurationProperty",
		reflect.TypeOf((*CfnRetrieverPropsMixin_NativeIndexConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnRetrieverPropsMixin.RetrieverConfigurationProperty",
		reflect.TypeOf((*CfnRetrieverPropsMixin_RetrieverConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperienceMixinProps",
		reflect.TypeOf((*CfnWebExperienceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperiencePropsMixin",
		reflect.TypeOf((*CfnWebExperiencePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWebExperiencePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperiencePropsMixin.BrowserExtensionConfigurationProperty",
		reflect.TypeOf((*CfnWebExperiencePropsMixin_BrowserExtensionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperiencePropsMixin.CustomizationConfigurationProperty",
		reflect.TypeOf((*CfnWebExperiencePropsMixin_CustomizationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperiencePropsMixin.IdentityProviderConfigurationProperty",
		reflect.TypeOf((*CfnWebExperiencePropsMixin_IdentityProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperiencePropsMixin.OpenIDConnectProviderConfigurationProperty",
		reflect.TypeOf((*CfnWebExperiencePropsMixin_OpenIDConnectProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnWebExperiencePropsMixin.SamlProviderConfigurationProperty",
		reflect.TypeOf((*CfnWebExperiencePropsMixin_SamlProviderConfigurationProperty)(nil)).Elem(),
	)
}
