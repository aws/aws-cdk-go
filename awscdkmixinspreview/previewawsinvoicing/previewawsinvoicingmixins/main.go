package previewawsinvoicingmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_invoicing.mixins.CfnInvoiceUnitMixinProps",
		reflect.TypeOf((*CfnInvoiceUnitMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_invoicing.mixins.CfnInvoiceUnitPropsMixin",
		reflect.TypeOf((*CfnInvoiceUnitPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInvoiceUnitPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_invoicing.mixins.CfnInvoiceUnitPropsMixin.ResourceTagProperty",
		reflect.TypeOf((*CfnInvoiceUnitPropsMixin_ResourceTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_invoicing.mixins.CfnInvoiceUnitPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnInvoiceUnitPropsMixin_RuleProperty)(nil)).Elem(),
	)
}
