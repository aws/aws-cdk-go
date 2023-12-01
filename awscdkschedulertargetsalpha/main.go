// The CDK Construct Library for Amazon Scheduler Targets
package awscdkschedulertargetsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-targets-alpha.CodeBuildStartBuild",
		reflect.TypeOf((*CodeBuildStartBuild)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_CodeBuildStartBuild{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awscdkscheduleralphaIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-targets-alpha.EventBridgePutEvents",
		reflect.TypeOf((*EventBridgePutEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_EventBridgePutEvents{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awscdkscheduleralphaIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-scheduler-targets-alpha.EventBridgePutEventsEntry",
		reflect.TypeOf((*EventBridgePutEventsEntry)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-targets-alpha.InspectorStartAssessmentRun",
		reflect.TypeOf((*InspectorStartAssessmentRun)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_InspectorStartAssessmentRun{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awscdkscheduleralphaIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-targets-alpha.LambdaInvoke",
		reflect.TypeOf((*LambdaInvoke)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaInvoke{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awscdkscheduleralphaIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-targets-alpha.ScheduleTargetBase",
		reflect.TypeOf((*ScheduleTargetBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			return &jsiiProxy_ScheduleTargetBase{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-scheduler-targets-alpha.ScheduleTargetBaseProps",
		reflect.TypeOf((*ScheduleTargetBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-targets-alpha.SnsPublish",
		reflect.TypeOf((*SnsPublish)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_SnsPublish{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awscdkscheduleralphaIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-targets-alpha.SqsSendMessage",
		reflect.TypeOf((*SqsSendMessage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_SqsSendMessage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awscdkscheduleralphaIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-scheduler-targets-alpha.SqsSendMessageProps",
		reflect.TypeOf((*SqsSendMessageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-targets-alpha.StepFunctionsStartExecution",
		reflect.TypeOf((*StepFunctionsStartExecution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_StepFunctionsStartExecution{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awscdkscheduleralphaIScheduleTarget)
			return &j
		},
	)
}
