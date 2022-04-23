package awsioteventsactions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_iotevents_actions.LambdaInvokeAction",
		reflect.TypeOf((*LambdaInvokeAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaInvokeAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsioteventsIAction)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iotevents_actions.SetVariableAction",
		reflect.TypeOf((*SetVariableAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SetVariableAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsioteventsIAction)
			return &j
		},
	)
}
