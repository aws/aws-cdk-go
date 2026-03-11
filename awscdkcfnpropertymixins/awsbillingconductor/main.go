package awsbillingconductor

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnBillingGroupMixinProps",
		reflect.TypeOf((*CfnBillingGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnBillingGroupPropsMixin",
		reflect.TypeOf((*CfnBillingGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBillingGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnBillingGroupPropsMixin.AccountGroupingProperty",
		reflect.TypeOf((*CfnBillingGroupPropsMixin_AccountGroupingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnBillingGroupPropsMixin.ComputationPreferenceProperty",
		reflect.TypeOf((*CfnBillingGroupPropsMixin_ComputationPreferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnCustomLineItemMixinProps",
		reflect.TypeOf((*CfnCustomLineItemMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnCustomLineItemPropsMixin",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomLineItemPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnCustomLineItemPropsMixin.BillingPeriodRangeProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_BillingPeriodRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnCustomLineItemPropsMixin.CustomLineItemChargeDetailsProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_CustomLineItemChargeDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnCustomLineItemPropsMixin.CustomLineItemFlatChargeDetailsProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_CustomLineItemFlatChargeDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnCustomLineItemPropsMixin.CustomLineItemPercentageChargeDetailsProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_CustomLineItemPercentageChargeDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnCustomLineItemPropsMixin.LineItemFilterProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_LineItemFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnCustomLineItemPropsMixin.PresentationDetailsProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_PresentationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnPricingPlanMixinProps",
		reflect.TypeOf((*CfnPricingPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnPricingPlanPropsMixin",
		reflect.TypeOf((*CfnPricingPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPricingPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnPricingRuleMixinProps",
		reflect.TypeOf((*CfnPricingRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnPricingRulePropsMixin",
		reflect.TypeOf((*CfnPricingRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPricingRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnPricingRulePropsMixin.FreeTierProperty",
		reflect.TypeOf((*CfnPricingRulePropsMixin_FreeTierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_billingconductor.CfnPricingRulePropsMixin.TieringProperty",
		reflect.TypeOf((*CfnPricingRulePropsMixin_TieringProperty)(nil)).Elem(),
	)
}
