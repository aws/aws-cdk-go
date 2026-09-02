package awsinvoicing

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnInvoiceUnitMixinProps",
		reflect.TypeOf((*CfnInvoiceUnitMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnInvoiceUnitPropsMixin",
		reflect.TypeOf((*CfnInvoiceUnitPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInvoiceUnitPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnInvoiceUnitPropsMixin.ResourceTagProperty",
		reflect.TypeOf((*CfnInvoiceUnitPropsMixin_ResourceTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnInvoiceUnitPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnInvoiceUnitPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferenceMixinProps",
		reflect.TypeOf((*CfnProcurementPortalPreferenceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferencePropsMixin",
		reflect.TypeOf((*CfnProcurementPortalPreferencePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProcurementPortalPreferencePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferencePropsMixin.ContactProperty",
		reflect.TypeOf((*CfnProcurementPortalPreferencePropsMixin_ContactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferencePropsMixin.EinvoiceDeliveryPreferenceProperty",
		reflect.TypeOf((*CfnProcurementPortalPreferencePropsMixin_EinvoiceDeliveryPreferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferencePropsMixin.ProcurementPortalPreferenceSelectorProperty",
		reflect.TypeOf((*CfnProcurementPortalPreferencePropsMixin_ProcurementPortalPreferenceSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferencePropsMixin.PurchaseOrderDataSourceProperty",
		reflect.TypeOf((*CfnProcurementPortalPreferencePropsMixin_PurchaseOrderDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferencePropsMixin.TestEnvPreferenceProperty",
		reflect.TypeOf((*CfnProcurementPortalPreferencePropsMixin_TestEnvPreferenceProperty)(nil)).Elem(),
	)
}
