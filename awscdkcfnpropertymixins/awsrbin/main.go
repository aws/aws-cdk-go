package awsrbin

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rbin.CfnRuleMixinProps",
		reflect.TypeOf((*CfnRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_rbin.CfnRulePropsMixin",
		reflect.TypeOf((*CfnRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rbin.CfnRulePropsMixin.ResourceTagProperty",
		reflect.TypeOf((*CfnRulePropsMixin_ResourceTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rbin.CfnRulePropsMixin.RetentionPeriodProperty",
		reflect.TypeOf((*CfnRulePropsMixin_RetentionPeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rbin.CfnRulePropsMixin.UnlockDelayProperty",
		reflect.TypeOf((*CfnRulePropsMixin_UnlockDelayProperty)(nil)).Elem(),
	)
}
