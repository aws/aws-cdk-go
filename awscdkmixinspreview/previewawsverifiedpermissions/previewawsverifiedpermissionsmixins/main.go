package previewawsverifiedpermissionsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourceMixinProps",
		reflect.TypeOf((*CfnIdentitySourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentitySourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin.CognitoGroupConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_CognitoGroupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin.CognitoUserPoolConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_CognitoUserPoolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin.IdentitySourceConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_IdentitySourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin.IdentitySourceDetailsProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_IdentitySourceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin.OpenIdConnectAccessTokenConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_OpenIdConnectAccessTokenConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin.OpenIdConnectConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_OpenIdConnectConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin.OpenIdConnectGroupConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_OpenIdConnectGroupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin.OpenIdConnectIdentityTokenConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_OpenIdConnectIdentityTokenConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin.OpenIdConnectTokenSelectionProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_OpenIdConnectTokenSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyMixinProps",
		reflect.TypeOf((*CfnPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyPropsMixin",
		reflect.TypeOf((*CfnPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyPropsMixin.EntityIdentifierProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_EntityIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyPropsMixin.PolicyDefinitionProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_PolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyPropsMixin.StaticPolicyDefinitionProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_StaticPolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyPropsMixin.TemplateLinkedPolicyDefinitionProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_TemplateLinkedPolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyStoreMixinProps",
		reflect.TypeOf((*CfnPolicyStoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyStorePropsMixin",
		reflect.TypeOf((*CfnPolicyStorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyStorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyStorePropsMixin.DeletionProtectionProperty",
		reflect.TypeOf((*CfnPolicyStorePropsMixin_DeletionProtectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyStorePropsMixin.SchemaDefinitionProperty",
		reflect.TypeOf((*CfnPolicyStorePropsMixin_SchemaDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyStorePropsMixin.ValidationSettingsProperty",
		reflect.TypeOf((*CfnPolicyStorePropsMixin_ValidationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyTemplateMixinProps",
		reflect.TypeOf((*CfnPolicyTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyTemplatePropsMixin",
		reflect.TypeOf((*CfnPolicyTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
