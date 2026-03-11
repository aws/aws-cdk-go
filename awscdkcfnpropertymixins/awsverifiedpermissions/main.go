package awsverifiedpermissions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourceMixinProps",
		reflect.TypeOf((*CfnIdentitySourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourcePropsMixin",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentitySourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourcePropsMixin.CognitoGroupConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_CognitoGroupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourcePropsMixin.CognitoUserPoolConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_CognitoUserPoolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourcePropsMixin.IdentitySourceConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_IdentitySourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourcePropsMixin.IdentitySourceDetailsProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_IdentitySourceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourcePropsMixin.OpenIdConnectAccessTokenConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_OpenIdConnectAccessTokenConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourcePropsMixin.OpenIdConnectConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_OpenIdConnectConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourcePropsMixin.OpenIdConnectGroupConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_OpenIdConnectGroupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourcePropsMixin.OpenIdConnectIdentityTokenConfigurationProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_OpenIdConnectIdentityTokenConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnIdentitySourcePropsMixin.OpenIdConnectTokenSelectionProperty",
		reflect.TypeOf((*CfnIdentitySourcePropsMixin_OpenIdConnectTokenSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyMixinProps",
		reflect.TypeOf((*CfnPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyPropsMixin",
		reflect.TypeOf((*CfnPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyPropsMixin.EntityIdentifierProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_EntityIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyPropsMixin.PolicyDefinitionProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_PolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyPropsMixin.StaticPolicyDefinitionProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_StaticPolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyPropsMixin.TemplateLinkedPolicyDefinitionProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_TemplateLinkedPolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStoreMixinProps",
		reflect.TypeOf((*CfnPolicyStoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStorePropsMixin",
		reflect.TypeOf((*CfnPolicyStorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyStorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStorePropsMixin.DeletionProtectionProperty",
		reflect.TypeOf((*CfnPolicyStorePropsMixin_DeletionProtectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStorePropsMixin.EncryptionSettingsProperty",
		reflect.TypeOf((*CfnPolicyStorePropsMixin_EncryptionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStorePropsMixin.EncryptionStateProperty",
		reflect.TypeOf((*CfnPolicyStorePropsMixin_EncryptionStateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStorePropsMixin.KmsEncryptionSettingsProperty",
		reflect.TypeOf((*CfnPolicyStorePropsMixin_KmsEncryptionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStorePropsMixin.KmsEncryptionStateProperty",
		reflect.TypeOf((*CfnPolicyStorePropsMixin_KmsEncryptionStateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStorePropsMixin.SchemaDefinitionProperty",
		reflect.TypeOf((*CfnPolicyStorePropsMixin_SchemaDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStorePropsMixin.ValidationSettingsProperty",
		reflect.TypeOf((*CfnPolicyStorePropsMixin_ValidationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyTemplateMixinProps",
		reflect.TypeOf((*CfnPolicyTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyTemplatePropsMixin",
		reflect.TypeOf((*CfnPolicyTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
