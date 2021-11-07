package awsiotactions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.CloudWatchLogsAction",
		reflect.TypeOf((*CloudWatchLogsAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_CloudWatchLogsAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.CloudWatchLogsActionProps",
		reflect.TypeOf((*CloudWatchLogsActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.LambdaFunctionAction",
		reflect.TypeOf((*LambdaFunctionAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaFunctionAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
}
