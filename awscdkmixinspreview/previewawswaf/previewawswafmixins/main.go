package previewawswafmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnByteMatchSetMixinProps",
		reflect.TypeOf((*CfnByteMatchSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnByteMatchSetPropsMixin",
		reflect.TypeOf((*CfnByteMatchSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnByteMatchSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnByteMatchSetPropsMixin.ByteMatchTupleProperty",
		reflect.TypeOf((*CfnByteMatchSetPropsMixin_ByteMatchTupleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnByteMatchSetPropsMixin.FieldToMatchProperty",
		reflect.TypeOf((*CfnByteMatchSetPropsMixin_FieldToMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnIPSetMixinProps",
		reflect.TypeOf((*CfnIPSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnIPSetPropsMixin",
		reflect.TypeOf((*CfnIPSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIPSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnIPSetPropsMixin.IPSetDescriptorProperty",
		reflect.TypeOf((*CfnIPSetPropsMixin_IPSetDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnRuleMixinProps",
		reflect.TypeOf((*CfnRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnRulePropsMixin",
		reflect.TypeOf((*CfnRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnRulePropsMixin.PredicateProperty",
		reflect.TypeOf((*CfnRulePropsMixin_PredicateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSizeConstraintSetMixinProps",
		reflect.TypeOf((*CfnSizeConstraintSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSizeConstraintSetPropsMixin",
		reflect.TypeOf((*CfnSizeConstraintSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSizeConstraintSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSizeConstraintSetPropsMixin.FieldToMatchProperty",
		reflect.TypeOf((*CfnSizeConstraintSetPropsMixin_FieldToMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSizeConstraintSetPropsMixin.SizeConstraintProperty",
		reflect.TypeOf((*CfnSizeConstraintSetPropsMixin_SizeConstraintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSqlInjectionMatchSetMixinProps",
		reflect.TypeOf((*CfnSqlInjectionMatchSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSqlInjectionMatchSetPropsMixin",
		reflect.TypeOf((*CfnSqlInjectionMatchSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSqlInjectionMatchSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSqlInjectionMatchSetPropsMixin.FieldToMatchProperty",
		reflect.TypeOf((*CfnSqlInjectionMatchSetPropsMixin_FieldToMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSqlInjectionMatchSetPropsMixin.SqlInjectionMatchTupleProperty",
		reflect.TypeOf((*CfnSqlInjectionMatchSetPropsMixin_SqlInjectionMatchTupleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnWebACLMixinProps",
		reflect.TypeOf((*CfnWebACLMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnWebACLPropsMixin",
		reflect.TypeOf((*CfnWebACLPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWebACLPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnWebACLPropsMixin.ActivatedRuleProperty",
		reflect.TypeOf((*CfnWebACLPropsMixin_ActivatedRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnWebACLPropsMixin.WafActionProperty",
		reflect.TypeOf((*CfnWebACLPropsMixin_WafActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnXssMatchSetMixinProps",
		reflect.TypeOf((*CfnXssMatchSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnXssMatchSetPropsMixin",
		reflect.TypeOf((*CfnXssMatchSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnXssMatchSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnXssMatchSetPropsMixin.FieldToMatchProperty",
		reflect.TypeOf((*CfnXssMatchSetPropsMixin_FieldToMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnXssMatchSetPropsMixin.XssMatchTupleProperty",
		reflect.TypeOf((*CfnXssMatchSetPropsMixin_XssMatchTupleProperty)(nil)).Elem(),
	)
}
