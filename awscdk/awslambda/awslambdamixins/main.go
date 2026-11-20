package awslambdamixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda.mixins.FunctionRuntimeConfig",
		reflect.TypeOf((*FunctionRuntimeConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_FunctionRuntimeConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda.mixins.FunctionRuntimeConfigProps",
		reflect.TypeOf((*FunctionRuntimeConfigProps)(nil)).Elem(),
	)
}
