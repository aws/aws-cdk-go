package awskendraranking

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kendraranking.CfnExecutionPlanMixinProps",
		reflect.TypeOf((*CfnExecutionPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_kendraranking.CfnExecutionPlanPropsMixin",
		reflect.TypeOf((*CfnExecutionPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExecutionPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kendraranking.CfnExecutionPlanPropsMixin.CapacityUnitsConfigurationProperty",
		reflect.TypeOf((*CfnExecutionPlanPropsMixin_CapacityUnitsConfigurationProperty)(nil)).Elem(),
	)
}
