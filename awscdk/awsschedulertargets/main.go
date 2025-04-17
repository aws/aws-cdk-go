package awsschedulertargets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.CodeBuildStartBuild",
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
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.CodePipelineStartPipelineExecution",
		reflect.TypeOf((*CodePipelineStartPipelineExecution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_CodePipelineStartPipelineExecution{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler_targets.Ec2TaskProps",
		reflect.TypeOf((*Ec2TaskProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.EcsRunEc2Task",
		reflect.TypeOf((*EcsRunEc2Task)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_EcsRunEc2Task{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EcsRunTask)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.EcsRunFargateTask",
		reflect.TypeOf((*EcsRunFargateTask)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_EcsRunFargateTask{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EcsRunTask)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.EcsRunTask",
		reflect.TypeOf((*EcsRunTask)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_EcsRunTask{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler_targets.EcsRunTaskBaseProps",
		reflect.TypeOf((*EcsRunTaskBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.EventBridgePutEvents",
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
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler_targets.EventBridgePutEventsEntry",
		reflect.TypeOf((*EventBridgePutEventsEntry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler_targets.FargateTaskProps",
		reflect.TypeOf((*FargateTaskProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.FirehosePutRecord",
		reflect.TypeOf((*FirehosePutRecord)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_FirehosePutRecord{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.InspectorStartAssessmentRun",
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
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.KinesisStreamPutRecord",
		reflect.TypeOf((*KinesisStreamPutRecord)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisStreamPutRecord{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler_targets.KinesisStreamPutRecordProps",
		reflect.TypeOf((*KinesisStreamPutRecordProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.LambdaInvoke",
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
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler_targets.SageMakerPipelineParameter",
		reflect.TypeOf((*SageMakerPipelineParameter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.SageMakerStartPipelineExecution",
		reflect.TypeOf((*SageMakerStartPipelineExecution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_SageMakerStartPipelineExecution{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler_targets.SageMakerStartPipelineExecutionProps",
		reflect.TypeOf((*SageMakerStartPipelineExecutionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.ScheduleTargetBase",
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
		"aws-cdk-lib.aws_scheduler_targets.ScheduleTargetBaseProps",
		reflect.TypeOf((*ScheduleTargetBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.SnsPublish",
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
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.SqsSendMessage",
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
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler_targets.SqsSendMessageProps",
		reflect.TypeOf((*SqsSendMessageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.StepFunctionsStartExecution",
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
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler_targets.Tag",
		reflect.TypeOf((*Tag)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler_targets.Universal",
		reflect.TypeOf((*Universal)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTargetActionToRole", GoMethod: "AddTargetActionToRole"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindBaseTargetConfig", GoMethod: "BindBaseTargetConfig"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			j := jsiiProxy_Universal{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScheduleTargetBase)
			_jsii_.InitJsiiProxy(&j.Type__awsschedulerIScheduleTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler_targets.UniversalTargetProps",
		reflect.TypeOf((*UniversalTargetProps)(nil)).Elem(),
	)
}
