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
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetGroupMixinProps",
		reflect.TypeOf((*CfnLicenseAssetGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetGroupPropsMixin",
		reflect.TypeOf((*CfnLicenseAssetGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLicenseAssetGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetGroupPropsMixin.LicenseAssetGroupConfigurationProperty",
		reflect.TypeOf((*CfnLicenseAssetGroupPropsMixin_LicenseAssetGroupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseAssetGroupPropsMixin.LicenseAssetGroupPropertyProperty",
		reflect.TypeOf((*CfnLicenseAssetGroupPropsMixin_LicenseAssetGroupPropertyProperty)(nil)).Elem(),
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
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseConfigurationMixinProps",
		reflect.TypeOf((*CfnLicenseConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseConfigurationPropsMixin",
		reflect.TypeOf((*CfnLicenseConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLicenseConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseConfigurationPropsMixin.ProductInformationFilterProperty",
		reflect.TypeOf((*CfnLicenseConfigurationPropsMixin_ProductInformationFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnLicenseConfigurationPropsMixin.ProductInformationProperty",
		reflect.TypeOf((*CfnLicenseConfigurationPropsMixin_ProductInformationProperty)(nil)).Elem(),
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
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnReportGeneratorMixinProps",
		reflect.TypeOf((*CfnReportGeneratorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnReportGeneratorPropsMixin",
		reflect.TypeOf((*CfnReportGeneratorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReportGeneratorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnReportGeneratorPropsMixin.ReportContextProperty",
		reflect.TypeOf((*CfnReportGeneratorPropsMixin_ReportContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnReportGeneratorPropsMixin.ReportFrequencyProperty",
		reflect.TypeOf((*CfnReportGeneratorPropsMixin_ReportFrequencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanager.CfnReportGeneratorPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnReportGeneratorPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
}
