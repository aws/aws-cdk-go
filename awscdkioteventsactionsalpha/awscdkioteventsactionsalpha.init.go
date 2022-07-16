package awscdkioteventsactionsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iotevents-actions-alpha.LambdaInvokeAction",
		reflect.TypeOf((*LambdaInvokeAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaInvokeAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkioteventsalphaIAction)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iotevents-actions-alpha.SetVariableAction",
		reflect.TypeOf((*SetVariableAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SetVariableAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkioteventsalphaIAction)
			return &j
		},
	)
}
