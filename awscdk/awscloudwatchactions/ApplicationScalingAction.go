package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions/internal"
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
//   var stepScalingAction StepScalingAction
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

	if err := validateNewApplicationScalingActionParameters(stepScalingAction); err != nil {
		panic(err)
	}
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
	if err := a.validateBindParameters(_scope, _alarm); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope, _alarm},
		&returns,
	)

	return returns
}

