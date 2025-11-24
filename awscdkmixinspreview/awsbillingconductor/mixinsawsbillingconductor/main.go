package mixinsawsbillingconductor

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnBillingGroupMixinProps",
		reflect.TypeOf((*CfnBillingGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnBillingGroupPropsMixin",
		reflect.TypeOf((*CfnBillingGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBillingGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnBillingGroupPropsMixin.AccountGroupingProperty",
		reflect.TypeOf((*CfnBillingGroupPropsMixin_AccountGroupingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnBillingGroupPropsMixin.ComputationPreferenceProperty",
		reflect.TypeOf((*CfnBillingGroupPropsMixin_ComputationPreferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnCustomLineItemMixinProps",
		reflect.TypeOf((*CfnCustomLineItemMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnCustomLineItemPropsMixin",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomLineItemPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnCustomLineItemPropsMixin.BillingPeriodRangeProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_BillingPeriodRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnCustomLineItemPropsMixin.CustomLineItemChargeDetailsProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_CustomLineItemChargeDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnCustomLineItemPropsMixin.CustomLineItemFlatChargeDetailsProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_CustomLineItemFlatChargeDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnCustomLineItemPropsMixin.CustomLineItemPercentageChargeDetailsProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_CustomLineItemPercentageChargeDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnCustomLineItemPropsMixin.LineItemFilterProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_LineItemFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnCustomLineItemPropsMixin.PresentationDetailsProperty",
		reflect.TypeOf((*CfnCustomLineItemPropsMixin_PresentationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnPricingPlanMixinProps",
		reflect.TypeOf((*CfnPricingPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnPricingPlanPropsMixin",
		reflect.TypeOf((*CfnPricingPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPricingPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnPricingRuleMixinProps",
		reflect.TypeOf((*CfnPricingRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnPricingRulePropsMixin",
		reflect.TypeOf((*CfnPricingRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPricingRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnPricingRulePropsMixin.FreeTierProperty",
		reflect.TypeOf((*CfnPricingRulePropsMixin_FreeTierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnPricingRulePropsMixin.TieringProperty",
		reflect.TypeOf((*CfnPricingRulePropsMixin_TieringProperty)(nil)).Elem(),
	)
}
