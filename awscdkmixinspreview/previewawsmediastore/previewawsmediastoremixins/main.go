package previewawsmediastoremixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediastore.mixins.CfnContainerMixinProps",
		reflect.TypeOf((*CfnContainerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediastore.mixins.CfnContainerPropsMixin",
		reflect.TypeOf((*CfnContainerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContainerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediastore.mixins.CfnContainerPropsMixin.CorsRuleProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_CorsRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediastore.mixins.CfnContainerPropsMixin.MetricPolicyProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_MetricPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediastore.mixins.CfnContainerPropsMixin.MetricPolicyRuleProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_MetricPolicyRuleProperty)(nil)).Elem(),
	)
}
