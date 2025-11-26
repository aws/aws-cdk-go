package previewawsdatapipelinemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelineMixinProps",
		reflect.TypeOf((*CfnPipelineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin",
		reflect.TypeOf((*CfnPipelinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipelinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin.FieldProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_FieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin.ParameterAttributeProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ParameterAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin.ParameterObjectProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ParameterObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin.ParameterValueProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ParameterValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin.PipelineObjectProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_PipelineObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin.PipelineTagProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_PipelineTagProperty)(nil)).Elem(),
	)
}
