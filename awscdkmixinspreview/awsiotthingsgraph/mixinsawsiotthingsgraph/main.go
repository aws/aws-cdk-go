package mixinsawsiotthingsgraph

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotthingsgraph.mixins.CfnFlowTemplateMixinProps",
		reflect.TypeOf((*CfnFlowTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotthingsgraph.mixins.CfnFlowTemplatePropsMixin",
		reflect.TypeOf((*CfnFlowTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotthingsgraph.mixins.CfnFlowTemplatePropsMixin.DefinitionDocumentProperty",
		reflect.TypeOf((*CfnFlowTemplatePropsMixin_DefinitionDocumentProperty)(nil)).Elem(),
	)
}
