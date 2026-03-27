package awselementalinference

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elementalinference.CfnFeedMixinProps",
		reflect.TypeOf((*CfnFeedMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elementalinference.CfnFeedPropsMixin",
		reflect.TypeOf((*CfnFeedPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFeedPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elementalinference.CfnFeedPropsMixin.ClippingConfigProperty",
		reflect.TypeOf((*CfnFeedPropsMixin_ClippingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elementalinference.CfnFeedPropsMixin.GetOutputProperty",
		reflect.TypeOf((*CfnFeedPropsMixin_GetOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elementalinference.CfnFeedPropsMixin.OutputConfigProperty",
		reflect.TypeOf((*CfnFeedPropsMixin_OutputConfigProperty)(nil)).Elem(),
	)
}
