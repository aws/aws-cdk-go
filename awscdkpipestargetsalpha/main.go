// The CDK Construct Library for Amazon EventBridge Pipes Targets
package awscdkpipestargetsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-targets-alpha.SfnStateMachine",
		reflect.TypeOf((*SfnStateMachine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_SfnStateMachine{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaITarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-targets-alpha.SfnStateMachineParameters",
		reflect.TypeOf((*SfnStateMachineParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-targets-alpha.SqsTarget",
		reflect.TypeOf((*SqsTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_SqsTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaITarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-targets-alpha.SqsTargetParameters",
		reflect.TypeOf((*SqsTargetParameters)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-pipes-targets-alpha.StateMachineInvocationType",
		reflect.TypeOf((*StateMachineInvocationType)(nil)).Elem(),
		map[string]interface{}{
			"FIRE_AND_FORGET": StateMachineInvocationType_FIRE_AND_FORGET,
			"REQUEST_RESPONSE": StateMachineInvocationType_REQUEST_RESPONSE,
		},
	)
}
