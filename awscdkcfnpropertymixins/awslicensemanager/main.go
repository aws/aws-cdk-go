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
