package awslambdadestinations

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_destinations.EventBridgeDestination",
		reflect.TypeOf((*EventBridgeDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_EventBridgeDestination{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIDestination)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_destinations.LambdaDestination",
		reflect.TypeOf((*LambdaDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaDestination{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_destinations.LambdaDestinationOptions",
		reflect.TypeOf((*LambdaDestinationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_destinations.SnsDestination",
		reflect.TypeOf((*SnsDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SnsDestination{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIDestination)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_destinations.SqsDestination",
		reflect.TypeOf((*SqsDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SqsDestination{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIDestination)
			return &j
		},
	)
}
