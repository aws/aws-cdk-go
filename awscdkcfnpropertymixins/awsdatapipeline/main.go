package awsdatapipeline

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datapipeline.CfnPipelineMixinProps",
		reflect.TypeOf((*CfnPipelineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datapipeline.CfnPipelinePropsMixin",
		reflect.TypeOf((*CfnPipelinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipelinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datapipeline.CfnPipelinePropsMixin.FieldProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_FieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datapipeline.CfnPipelinePropsMixin.ParameterAttributeProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ParameterAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datapipeline.CfnPipelinePropsMixin.ParameterObjectProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ParameterObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datapipeline.CfnPipelinePropsMixin.ParameterValueProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ParameterValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datapipeline.CfnPipelinePropsMixin.PipelineObjectProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_PipelineObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datapipeline.CfnPipelinePropsMixin.PipelineTagProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_PipelineTagProperty)(nil)).Elem(),
	)
}
