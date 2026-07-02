package awscognito

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolMixinProps",
		reflect.TypeOf((*CfnIdentityPoolMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolPrincipalTagMixinProps",
		reflect.TypeOf((*CfnIdentityPoolPrincipalTagMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolPrincipalTagPropsMixin",
		reflect.TypeOf((*CfnIdentityPoolPrincipalTagPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentityPoolPrincipalTagPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolPropsMixin",
		reflect.TypeOf((*CfnIdentityPoolPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentityPoolPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolPropsMixin.CognitoIdentityProviderProperty",
		reflect.TypeOf((*CfnIdentityPoolPropsMixin_CognitoIdentityProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolPropsMixin.CognitoStreamsProperty",
		reflect.TypeOf((*CfnIdentityPoolPropsMixin_CognitoStreamsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolPropsMixin.PushSyncProperty",
		reflect.TypeOf((*CfnIdentityPoolPropsMixin_PushSyncProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolRoleAttachmentMixinProps",
		reflect.TypeOf((*CfnIdentityPoolRoleAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolRoleAttachmentPropsMixin",
		reflect.TypeOf((*CfnIdentityPoolRoleAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentityPoolRoleAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolRoleAttachmentPropsMixin.MappingRuleProperty",
		reflect.TypeOf((*CfnIdentityPoolRoleAttachmentPropsMixin_MappingRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolRoleAttachmentPropsMixin.RoleMappingProperty",
		reflect.TypeOf((*CfnIdentityPoolRoleAttachmentPropsMixin_RoleMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolRoleAttachmentPropsMixin.RulesConfigurationTypeProperty",
		reflect.TypeOf((*CfnIdentityPoolRoleAttachmentPropsMixin_RulesConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnLogDeliveryConfigurationMixinProps",
		reflect.TypeOf((*CfnLogDeliveryConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnLogDeliveryConfigurationPropsMixin",
		reflect.TypeOf((*CfnLogDeliveryConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogDeliveryConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnLogDeliveryConfigurationPropsMixin.CloudWatchLogsConfigurationProperty",
		reflect.TypeOf((*CfnLogDeliveryConfigurationPropsMixin_CloudWatchLogsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnLogDeliveryConfigurationPropsMixin.FirehoseConfigurationProperty",
		reflect.TypeOf((*CfnLogDeliveryConfigurationPropsMixin_FirehoseConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnLogDeliveryConfigurationPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnLogDeliveryConfigurationPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnLogDeliveryConfigurationPropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnLogDeliveryConfigurationPropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnManagedLoginBrandingMixinProps",
		reflect.TypeOf((*CfnManagedLoginBrandingMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnManagedLoginBrandingPropsMixin",
		reflect.TypeOf((*CfnManagedLoginBrandingPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnManagedLoginBrandingPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnManagedLoginBrandingPropsMixin.AssetTypeProperty",
		reflect.TypeOf((*CfnManagedLoginBrandingPropsMixin_AssetTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnTermsMixinProps",
		reflect.TypeOf((*CfnTermsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnTermsPropsMixin",
		reflect.TypeOf((*CfnTermsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTermsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolClientMixinProps",
		reflect.TypeOf((*CfnUserPoolClientMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolClientPropsMixin",
		reflect.TypeOf((*CfnUserPoolClientPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolClientPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolClientPropsMixin.AnalyticsConfigurationProperty",
		reflect.TypeOf((*CfnUserPoolClientPropsMixin_AnalyticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolClientPropsMixin.RefreshTokenRotationProperty",
		reflect.TypeOf((*CfnUserPoolClientPropsMixin_RefreshTokenRotationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolClientPropsMixin.TokenValidityUnitsProperty",
		reflect.TypeOf((*CfnUserPoolClientPropsMixin_TokenValidityUnitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolDomainMixinProps",
		reflect.TypeOf((*CfnUserPoolDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolDomainPropsMixin",
		reflect.TypeOf((*CfnUserPoolDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolDomainPropsMixin.CustomDomainConfigTypeProperty",
		reflect.TypeOf((*CfnUserPoolDomainPropsMixin_CustomDomainConfigTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolDomainPropsMixin.FailoverTypeProperty",
		reflect.TypeOf((*CfnUserPoolDomainPropsMixin_FailoverTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolDomainPropsMixin.RoutingTypeProperty",
		reflect.TypeOf((*CfnUserPoolDomainPropsMixin_RoutingTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolGroupMixinProps",
		reflect.TypeOf((*CfnUserPoolGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolGroupPropsMixin",
		reflect.TypeOf((*CfnUserPoolGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolIdentityProviderMixinProps",
		reflect.TypeOf((*CfnUserPoolIdentityProviderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolIdentityProviderPropsMixin",
		reflect.TypeOf((*CfnUserPoolIdentityProviderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolIdentityProviderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolMixinProps",
		reflect.TypeOf((*CfnUserPoolMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin",
		reflect.TypeOf((*CfnUserPoolPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.AccountRecoverySettingProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_AccountRecoverySettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.AdminCreateUserConfigProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_AdminCreateUserConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.AdvancedSecurityAdditionalFlowsProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_AdvancedSecurityAdditionalFlowsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.CustomEmailSenderProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_CustomEmailSenderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.CustomSMSSenderProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_CustomSMSSenderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.DeviceConfigurationProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_DeviceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.EmailConfigurationProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_EmailConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.InboundFederationProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_InboundFederationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.InviteMessageTemplateProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_InviteMessageTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.LambdaConfigProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_LambdaConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.NumberAttributeConstraintsProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_NumberAttributeConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.PasswordPolicyProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_PasswordPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.PoliciesProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_PoliciesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.PreTokenGenerationConfigProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_PreTokenGenerationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.RecoveryOptionProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_RecoveryOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.SchemaAttributeProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_SchemaAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.SignInPolicyProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_SignInPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.SmsConfigurationProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_SmsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.StringAttributeConstraintsProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_StringAttributeConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.UserAttributeUpdateSettingsProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_UserAttributeUpdateSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.UserPoolAddOnsProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_UserPoolAddOnsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.UsernameConfigurationProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_UsernameConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolPropsMixin.VerificationMessageTemplateProperty",
		reflect.TypeOf((*CfnUserPoolPropsMixin_VerificationMessageTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolResourceServerMixinProps",
		reflect.TypeOf((*CfnUserPoolResourceServerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolResourceServerPropsMixin",
		reflect.TypeOf((*CfnUserPoolResourceServerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolResourceServerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolResourceServerPropsMixin.ResourceServerScopeTypeProperty",
		reflect.TypeOf((*CfnUserPoolResourceServerPropsMixin_ResourceServerScopeTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolRiskConfigurationAttachmentMixinProps",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolRiskConfigurationAttachmentPropsMixin",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolRiskConfigurationAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolRiskConfigurationAttachmentPropsMixin.AccountTakeoverActionTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_AccountTakeoverActionTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolRiskConfigurationAttachmentPropsMixin.AccountTakeoverActionsTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_AccountTakeoverActionsTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolRiskConfigurationAttachmentPropsMixin.AccountTakeoverRiskConfigurationTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_AccountTakeoverRiskConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolRiskConfigurationAttachmentPropsMixin.CompromisedCredentialsActionsTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_CompromisedCredentialsActionsTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolRiskConfigurationAttachmentPropsMixin.CompromisedCredentialsRiskConfigurationTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_CompromisedCredentialsRiskConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolRiskConfigurationAttachmentPropsMixin.NotifyConfigurationTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_NotifyConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolRiskConfigurationAttachmentPropsMixin.NotifyEmailTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_NotifyEmailTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolRiskConfigurationAttachmentPropsMixin.RiskExceptionConfigurationTypeProperty",
		reflect.TypeOf((*CfnUserPoolRiskConfigurationAttachmentPropsMixin_RiskExceptionConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolUICustomizationAttachmentMixinProps",
		reflect.TypeOf((*CfnUserPoolUICustomizationAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolUICustomizationAttachmentPropsMixin",
		reflect.TypeOf((*CfnUserPoolUICustomizationAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolUICustomizationAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolUserMixinProps",
		reflect.TypeOf((*CfnUserPoolUserMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolUserPropsMixin",
		reflect.TypeOf((*CfnUserPoolUserPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolUserPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolUserPropsMixin.AttributeTypeProperty",
		reflect.TypeOf((*CfnUserPoolUserPropsMixin_AttributeTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolUserToGroupAttachmentMixinProps",
		reflect.TypeOf((*CfnUserPoolUserToGroupAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolUserToGroupAttachmentPropsMixin",
		reflect.TypeOf((*CfnUserPoolUserToGroupAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolUserToGroupAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
