// The CDK Construct Library for Amazon EventBridge Pipes Targets
package awscdkpipestargetsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-targets-alpha.ApiDestinationTarget",
		reflect.TypeOf((*ApiDestinationTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_ApiDestinationTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaITarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-targets-alpha.ApiDestinationTargetParameters",
		reflect.TypeOf((*ApiDestinationTargetParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-targets-alpha.CloudWatchLogsTarget",
		reflect.TypeOf((*CloudWatchLogsTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_CloudWatchLogsTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaITarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-targets-alpha.CloudWatchLogsTargetParameters",
		reflect.TypeOf((*CloudWatchLogsTargetParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-targets-alpha.EventBridgeTarget",
		reflect.TypeOf((*EventBridgeTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_EventBridgeTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaITarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-targets-alpha.EventBridgeTargetParameters",
		reflect.TypeOf((*EventBridgeTargetParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-targets-alpha.KinesisTarget",
		reflect.TypeOf((*KinesisTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaITarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-targets-alpha.KinesisTargetParameters",
		reflect.TypeOf((*KinesisTargetParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-targets-alpha.LambdaFunction",
		reflect.TypeOf((*LambdaFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaFunction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaITarget)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-pipes-targets-alpha.LambdaFunctionInvocationType",
		reflect.TypeOf((*LambdaFunctionInvocationType)(nil)).Elem(),
		map[string]interface{}{
			"FIRE_AND_FORGET": LambdaFunctionInvocationType_FIRE_AND_FORGET,
			"REQUEST_RESPONSE": LambdaFunctionInvocationType_REQUEST_RESPONSE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-targets-alpha.LambdaFunctionParameters",
		reflect.TypeOf((*LambdaFunctionParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-targets-alpha.SageMakerTarget",
		reflect.TypeOf((*SageMakerTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_SageMakerTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaITarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-targets-alpha.SageMakerTargetParameters",
		reflect.TypeOf((*SageMakerTargetParameters)(nil)).Elem(),
	)
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
