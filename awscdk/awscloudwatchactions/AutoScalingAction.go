package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an AutoScaling StepScalingAction as an Alarm Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stepScalingAction StepScalingAction
//
//   autoScalingAction := awscdk.Aws_cloudwatch_actions.NewAutoScalingAction(stepScalingAction)
//
type AutoScalingAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an AutoScaling StepScalingAction as an alarm action.
	Bind(scope constructs.Construct, alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for AutoScalingAction
type jsiiProxy_AutoScalingAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

func NewAutoScalingAction(stepScalingAction awsautoscaling.StepScalingAction) AutoScalingAction {
	_init_.Initialize()

	if err := validateNewAutoScalingActionParameters(stepScalingAction); err != nil {
		panic(err)
	}
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

func (a *jsiiProxy_AutoScalingAction) Bind(scope constructs.Construct, alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
	if err := a.validateBindParameters(scope, alarm); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.AlarmActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, alarm},
		&returns,
	)

	return returns
}

