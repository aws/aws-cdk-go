package awslicensemanager

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnGrantMixinProps",
		reflect.TypeOf((*CfnGrantMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnGrantPropsMixin",
		reflect.TypeOf((*CfnGrantPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGrantPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetMixinProps",
		reflect.TypeOf((*CfnLicenseAssetRuleSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin",
		reflect.TypeOf((*CfnLicenseAssetRuleSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLicenseAssetRuleSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin.AndRuleStatementProperty",
		reflect.TypeOf((*CfnLicenseAssetRuleSetPropsMixin_AndRuleStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin.InstanceRuleStatementProperty",
		reflect.TypeOf((*CfnLicenseAssetRuleSetPropsMixin_InstanceRuleStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin.LicenseAssetRuleProperty",
		reflect.TypeOf((*CfnLicenseAssetRuleSetPropsMixin_LicenseAssetRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin.LicenseConfigurationRuleStatementProperty",
		reflect.TypeOf((*CfnLicenseAssetRuleSetPropsMixin_LicenseConfigurationRuleStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin.LicenseRuleStatementProperty",
		reflect.TypeOf((*CfnLicenseAssetRuleSetPropsMixin_LicenseRuleStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin.MatchingRuleStatementProperty",
		reflect.TypeOf((*CfnLicenseAssetRuleSetPropsMixin_MatchingRuleStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin.OrRuleStatementProperty",
		reflect.TypeOf((*CfnLicenseAssetRuleSetPropsMixin_OrRuleStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetRuleSetPropsMixin.RuleStatementProperty",
		reflect.TypeOf((*CfnLicenseAssetRuleSetPropsMixin_RuleStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseMixinProps",
		reflect.TypeOf((*CfnLicenseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicensePropsMixin",
		reflect.TypeOf((*CfnLicensePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLicensePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicensePropsMixin.BorrowConfigurationProperty",
		reflect.TypeOf((*CfnLicensePropsMixin_BorrowConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicensePropsMixin.ConsumptionConfigurationProperty",
		reflect.TypeOf((*CfnLicensePropsMixin_ConsumptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicensePropsMixin.EntitlementProperty",
		reflect.TypeOf((*CfnLicensePropsMixin_EntitlementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicensePropsMixin.IssuerDataProperty",
		reflect.TypeOf((*CfnLicensePropsMixin_IssuerDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicensePropsMixin.MetadataProperty",
		reflect.TypeOf((*CfnLicensePropsMixin_MetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicensePropsMixin.ProvisionalConfigurationProperty",
		reflect.TypeOf((*CfnLicensePropsMixin_ProvisionalConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicensePropsMixin.ValidityDateFormatProperty",
		reflect.TypeOf((*CfnLicensePropsMixin_ValidityDateFormatProperty)(nil)).Elem(),
	)
}
