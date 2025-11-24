package mixinsawskendraranking

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kendraranking.mixins.CfnExecutionPlanMixinProps",
		reflect.TypeOf((*CfnExecutionPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kendraranking.mixins.CfnExecutionPlanPropsMixin",
		reflect.TypeOf((*CfnExecutionPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExecutionPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kendraranking.mixins.CfnExecutionPlanPropsMixin.CapacityUnitsConfigurationProperty",
		reflect.TypeOf((*CfnExecutionPlanPropsMixin_CapacityUnitsConfigurationProperty)(nil)).Elem(),
	)
}
