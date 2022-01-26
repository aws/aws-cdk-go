package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an ApplicationAutoScaling StepScalingAction as an Alarm Action.
//
// TODO: EXAMPLE
//
type ApplicationScalingAction interface {
	awscloudwatch.IAlarmAction
	Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for ApplicationScalingAction
type jsiiProxy_ApplicationScalingAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewApplicationScalingAction(stepScalingAction awsapplicationautoscaling.StepScalingAction) ApplicationScalingAction {
	_init_.Initialize()

	j := jsiiProxy_ApplicationScalingAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.ApplicationScalingAction",
		[]interface{}{stepScalingAction},
		&j,
	)

	return &j
}

func NewApplicationScalingAction_Override(a ApplicationScalingAction, stepScalingAction awsapplicationautoscaling.StepScalingAction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.ApplicationScalingAction",
		[]interface{}{stepScalingAction},
		a,
	)
}

// Returns an alarm action configuration to use an ApplicationScaling StepScalingAction as an alarm action.
func (a *jsiiProxy_ApplicationScalingAction) Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope, _alarm},
		&returns,
	)

	return returns
}

// Use an AutoScaling StepScalingAction as an Alarm Action.
//
// TODO: EXAMPLE
//
type AutoScalingAction interface {
	awscloudwatch.IAlarmAction
	Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for AutoScalingAction
type jsiiProxy_AutoScalingAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewAutoScalingAction(stepScalingAction awsautoscaling.StepScalingAction) AutoScalingAction {
	_init_.Initialize()

	j := jsiiProxy_AutoScalingAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.AutoScalingAction",
		[]interface{}{stepScalingAction},
		&j,
	)

	return &j
}

func NewAutoScalingAction_Override(a AutoScalingAction, stepScalingAction awsautoscaling.StepScalingAction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.AutoScalingAction",
		[]interface{}{stepScalingAction},
		a,
	)
}

// Returns an alarm action configuration to use an AutoScaling StepScalingAction as an alarm action.
func (a *jsiiProxy_AutoScalingAction) Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope, _alarm},
		&returns,
	)

	return returns
}

// Use an EC2 action as an Alarm action.
//
// TODO: EXAMPLE
//
type Ec2Action interface {
	awscloudwatch.IAlarmAction
	Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for Ec2Action
type jsiiProxy_Ec2Action struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewEc2Action(instanceAction Ec2InstanceAction) Ec2Action {
	_init_.Initialize()

	j := jsiiProxy_Ec2Action{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.Ec2Action",
		[]interface{}{instanceAction},
		&j,
	)

	return &j
}

func NewEc2Action_Override(e Ec2Action, instanceAction Ec2InstanceAction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.Ec2Action",
		[]interface{}{instanceAction},
		e,
	)
}

// Returns an alarm action configuration to use an EC2 action as an alarm action.
func (e *jsiiProxy_Ec2Action) Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope, _alarm},
		&returns,
	)

	return returns
}

// Types of EC2 actions available.
type Ec2InstanceAction string

const (
	Ec2InstanceAction_STOP Ec2InstanceAction = "STOP"
	Ec2InstanceAction_TERMINATE Ec2InstanceAction = "TERMINATE"
	Ec2InstanceAction_RECOVER Ec2InstanceAction = "RECOVER"
	Ec2InstanceAction_REBOOT Ec2InstanceAction = "REBOOT"
)

// Use an SNS topic as an alarm action.
//
// TODO: EXAMPLE
//
type SnsAction interface {
	awscloudwatch.IAlarmAction
	Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for SnsAction
type jsiiProxy_SnsAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewSnsAction(topic awssns.ITopic) SnsAction {
	_init_.Initialize()

	j := jsiiProxy_SnsAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.SnsAction",
		[]interface{}{topic},
		&j,
	)

	return &j
}

func NewSnsAction_Override(s SnsAction, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.SnsAction",
		[]interface{}{topic},
		s,
	)
}

// Returns an alarm action configuration to use an SNS topic as an alarm action.
func (s *jsiiProxy_SnsAction) Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _alarm},
		&returns,
	)

	return returns
}

