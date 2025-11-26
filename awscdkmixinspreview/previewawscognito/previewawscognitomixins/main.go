package previewawscognitomixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolMixinProps",
		reflect.TypeOf((*CfnIdentityPoolMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPrincipalTagMixinProps",
		reflect.TypeOf((*CfnIdentityPoolPrincipalTagMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPrincipalTagPropsMixin",
		reflect.TypeOf((*CfnIdentityPoolPrincipalTagPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentityPoolPrincipalTagPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPropsMixin",
		reflect.TypeOf((*CfnIdentityPoolPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentityPoolPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPropsMixin.CognitoIdentityProviderProperty",
		reflect.TypeOf((*CfnIdentityPoolPropsMixin_CognitoIdentityProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPropsMixin.CognitoStreamsProperty",
		reflect.TypeOf((*CfnIdentityPoolPropsMixin_CognitoStreamsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPropsMixin.PushSyncProperty",
		reflect.TypeOf((*CfnIdentityPoolPropsMixin_PushSyncProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolRoleAttachmentMixinProps",
		reflect.TypeOf((*CfnIdentityPoolRoleAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolRoleAttachmentPropsMixin",
		reflect.TypeOf((*CfnIdentityPoolRoleAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentityPoolRoleAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolRoleAttachmentPropsMixin.MappingRuleProperty",
		reflect.TypeOf((*CfnIdentityPoolRoleAttachmentPropsMixin_MappingRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolRoleAttachmentPropsMixin.RoleMappingProperty",
		reflect.TypeOf((*CfnIdentityPoolRoleAttachmentPropsMixin_RoleMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolRoleAttachmentPropsMixin.RulesConfigurationTypeProperty",
		reflect.TypeOf((*CfnIdentityPoolRoleAttachmentPropsMixin_RulesConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnLogDeliveryConfigurationMixinProps",
		reflect.TypeOf((*CfnLogDeliveryConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnLogDeliveryConfigurationPropsMixin",
		reflect.TypeOf((*CfnLogDeliveryConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogDeliveryConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnLogDeliveryConfigurationPropsMixin.CloudWatchLogsConfigurationProperty",
		reflect.TypeOf((*CfnLogDeliveryConfigurationPropsMixin_CloudWatchLogsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnLogDeliveryConfigurationPropsMixin.FirehoseConfigurationProperty",
		reflect.TypeOf((*CfnLogDeliveryConfigurationPropsMixin_FirehoseConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnLogDeliveryConfigurationPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnLogDeliveryConfigurationPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnLogDeliveryConfigurationPropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnLogDeliveryConfigurationPropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnManagedLoginBrandingMixinProps",
		reflect.TypeOf((*CfnManagedLoginBrandingMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnManagedLoginBrandingPropsMixin",
		reflect.TypeOf((*CfnManagedLoginBrandingPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnManagedLoginBrandingPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnManagedLoginBrandingPropsMixin.AssetTypeProperty",
		reflect.TypeOf((*CfnManagedLoginBrandingPropsMixin_AssetTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnTermsMixinProps",
		reflect.TypeOf((*CfnTermsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnTermsPropsMixin",
		reflect.TypeOf((*CfnTermsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTermsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogs",
		reflect.TypeOf((*CfnUserPoolApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnUserPoolApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolClientMixinProps",
		reflect.TypeOf((*CfnUserPoolClientMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolClientPropsMixin",
		reflect.TypeOf((*CfnUserPoolClientPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolClientPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolClientPropsMixin.AnalyticsConfigurationProperty",
		reflect.TypeOf((*CfnUserPoolClientPropsMixin_AnalyticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolClientPropsMixin.RefreshTokenRotationProperty",
		reflect.TypeOf((*CfnUserPoolClientPropsMixin_RefreshTokenRotationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolClientPropsMixin.TokenValidityUnitsProperty",
		reflect.TypeOf((*CfnUserPoolClientPropsMixin_TokenValidityUnitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolDomainMixinProps",
		reflect.TypeOf((*CfnUserPoolDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolDomainPropsMixin",
		reflect.TypeOf((*CfnUserPoolDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolDomainPropsMixin.CustomDomainConfigTypeProperty",
		reflect.TypeOf((*CfnUserPoolDomainPropsMixin_CustomDomainConfigTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolGroupMixinProps",
		reflect.TypeOf((*CfnUserPoolGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolGroupPropsMixin",
		reflect.TypeOf((*CfnUserPoolGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolIdentityProviderMixinProps",
		reflect.TypeOf((*CfnUserPoolIdentityProviderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolIdentityProviderPropsMixin",
		reflect.TypeOf((*CfnUserPoolIdentityProviderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolIdentityProviderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolLogsMixin",
		reflect.TypeOf((*CfnUserPoolLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolMixinProps",
		reflect.TypeOf((*CfnUserPoolMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin",
		reflect.TypeOf((*CfnUserPoolPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.AccountRecoverySettingProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_AccountRecoverySettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.AdminCreateUserConfigProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_AdminCreateUserConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.AdvancedSecurityAdditionalFlowsProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_AdvancedSecurityAdditionalFlowsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.CustomEmailSenderProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_CustomEmailSenderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.CustomSMSSenderProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_CustomSMSSenderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.DeviceConfigurationProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_DeviceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.EmailConfigurationProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_EmailConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.InviteMessageTemplateProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_InviteMessageTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.LambdaConfigProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_LambdaConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.NumberAttributeConstraintsProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_NumberAttributeConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.PasswordPolicyProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_PasswordPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.PoliciesProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_PoliciesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.PreTokenGenerationConfigProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_PreTokenGenerationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.RecoveryOptionProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_RecoveryOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.SchemaAttributeProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_SchemaAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.SignInPolicyProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_SignInPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.SmsConfigurationProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_SmsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.StringAttributeConstraintsProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_StringAttributeConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.UserAttributeUpdateSettingsProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_UserAttributeUpdateSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.UserPoolAddOnsProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_UserPoolAddOnsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.UsernameConfigurationProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_UsernameConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolPropsMixin.VerificationMessageTemplateProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_VerificationMessageTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolResourceServerMixinProps",
		reflect.TypeOf((*CfnUserPoolResourceServerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolResourceServerPropsMixin",
		reflect.TypeOf((*CfnUserPoolResourceServerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolResourceServerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolResourceServerPropsMixin.ResourceServerScopeTypeProperty",
		reflect.TypeOf((*CfnUserPoolResourceServerPropsMixin_ResourceServerScopeTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentMixinProps",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolRiskConfigurationAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin.AccountTakeoverActionTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_AccountTakeoverActionTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin.AccountTakeoverActionsTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_AccountTakeoverActionsTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin.AccountTakeoverRiskConfigurationTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_AccountTakeoverRiskConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin.CompromisedCredentialsActionsTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_CompromisedCredentialsActionsTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin.CompromisedCredentialsRiskConfigurationTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_CompromisedCredentialsRiskConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin.NotifyConfigurationTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_NotifyConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin.NotifyEmailTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_NotifyEmailTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin.RiskExceptionConfigurationTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_RiskExceptionConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolUICustomizationAttachmentMixinProps",
		reflect.TypeOf((*CfnUserPoolUICustomizationAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolUICustomizationAttachmentPropsMixin",
		reflect.TypeOf((*CfnUserPoolUICustomizationAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolUICustomizationAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolUserMixinProps",
		reflect.TypeOf((*CfnUserPoolUserMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolUserPropsMixin",
		reflect.TypeOf((*CfnUserPoolUserPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolUserPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolUserPropsMixin.AttributeTypeProperty",
		reflect.TypeOf((*CfnUserPoolUserPropsMixin_AttributeTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolUserToGroupAttachmentMixinProps",
		reflect.TypeOf((*CfnUserPoolUserToGroupAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolUserToGroupAttachmentPropsMixin",
		reflect.TypeOf((*CfnUserPoolUserToGroupAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolUserToGroupAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
