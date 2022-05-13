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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stepScalingAction stepScalingAction
//
//   applicationScalingAction := awscdk.Aws_cloudwatch_actions.NewApplicationScalingAction(stepScalingAction)
//
type ApplicationScalingAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an ApplicationScaling StepScalingAction as an alarm action.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stepScalingAction stepScalingAction
//
//   autoScalingAction := awscdk.Aws_cloudwatch_actions.NewAutoScalingAction(stepScalingAction)
//
type AutoScalingAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an AutoScaling StepScalingAction as an alarm action.
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
// Example:
//   // Alarm must be configured with an EC2 per-instance metric
//   var alarm alarm
//
//   // Attach a reboot when alarm triggers
//   alarm.addAlarmAction(
//   actions.NewEc2Action(actions.ec2InstanceAction_REBOOT))
//
type Ec2Action interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an EC2 action as an alarm action.
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
//
// Example:
//   // Alarm must be configured with an EC2 per-instance metric
//   var alarm alarm
//
//   // Attach a reboot when alarm triggers
//   alarm.addAlarmAction(
//   actions.NewEc2Action(actions.ec2InstanceAction_REBOOT))
//
type Ec2InstanceAction string

const (
	// Stop the instance.
	Ec2InstanceAction_STOP Ec2InstanceAction = "STOP"
	// Terminatethe instance.
	Ec2InstanceAction_TERMINATE Ec2InstanceAction = "TERMINATE"
	// Recover the instance.
	Ec2InstanceAction_RECOVER Ec2InstanceAction = "RECOVER"
	// Reboot the instance.
	Ec2InstanceAction_REBOOT Ec2InstanceAction = "REBOOT"
)

// Types of OpsItem category available.
//
// Example:
//   var alarm alarm
//
//   // Create an OpsItem with specific severity and category when alarm triggers
//   alarm.addAlarmAction(
//   actions.NewSsmAction(actions.opsItemSeverity_CRITICAL, actions.opsItemCategory_PERFORMANCE))
//
type OpsItemCategory string

const (
	// Set the category to availability.
	OpsItemCategory_AVAILABILITY OpsItemCategory = "AVAILABILITY"
	// Set the category to cost.
	OpsItemCategory_COST OpsItemCategory = "COST"
	// Set the category to performance.
	OpsItemCategory_PERFORMANCE OpsItemCategory = "PERFORMANCE"
	// Set the category to recovery.
	OpsItemCategory_RECOVERY OpsItemCategory = "RECOVERY"
	// Set the category to security.
	OpsItemCategory_SECURITY OpsItemCategory = "SECURITY"
)

// Types of OpsItem severity available.
//
// Example:
//   var alarm alarm
//
//   // Create an OpsItem with specific severity and category when alarm triggers
//   alarm.addAlarmAction(
//   actions.NewSsmAction(actions.opsItemSeverity_CRITICAL, actions.opsItemCategory_PERFORMANCE))
//
type OpsItemSeverity string

const (
	// Set the severity to critical.
	OpsItemSeverity_CRITICAL OpsItemSeverity = "CRITICAL"
	// Set the severity to high.
	OpsItemSeverity_HIGH OpsItemSeverity = "HIGH"
	// Set the severity to medium.
	OpsItemSeverity_MEDIUM OpsItemSeverity = "MEDIUM"
	// Set the severity to low.
	OpsItemSeverity_LOW OpsItemSeverity = "LOW"
)

// Use an SNS topic as an alarm action.
//
// Example:
//   import cw_actions "github.com/aws/aws-cdk-go/awscdk"
//   var alarm alarm
//
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//   alarm.addAlarmAction(cw_actions.NewSnsAction(topic))
//
type SnsAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an SNS topic as an alarm action.
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

// Use an SSM OpsItem action as an Alarm action.
//
// Example:
//   var alarm alarm
//
//   // Create an OpsItem with specific severity and category when alarm triggers
//   alarm.addAlarmAction(
//   actions.NewSsmAction(actions.opsItemSeverity_CRITICAL, actions.opsItemCategory_PERFORMANCE))
//
type SsmAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an SSM OpsItem action as an alarm action.
	Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for SsmAction
type jsiiProxy_SsmAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewSsmAction(severity OpsItemSeverity, category OpsItemCategory) SsmAction {
	_init_.Initialize()

	j := jsiiProxy_SsmAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.SsmAction",
		[]interface{}{severity, category},
		&j,
	)

	return &j
}

func NewSsmAction_Override(s SsmAction, severity OpsItemSeverity, category OpsItemCategory) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch_actions.SsmAction",
		[]interface{}{severity, category},
		s,
	)
}

func (s *jsiiProxy_SsmAction) Bind(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _alarm},
		&returns,
	)

	return returns
}

