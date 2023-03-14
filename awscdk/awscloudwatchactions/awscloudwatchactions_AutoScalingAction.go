package awscloudwatchactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatchactions/internal"
)

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
// Experimental.
type AutoScalingAction interface {
	awscloudwatch.IAlarmAction
	// Returns an alarm action configuration to use an AutoScaling StepScalingAction as an alarm action.
	// Experimental.
	Bind(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig
}

// The jsii proxy struct for AutoScalingAction
type jsiiProxy_AutoScalingAction struct {
	internal.Type__awscloudwatchIAlarmAction
}

// Experimental.
func NewAutoScalingAction(stepScalingAction awsautoscaling.StepScalingAction) AutoScalingAction {
	_init_.Initialize()

	if err := validateNewAutoScalingActionParameters(stepScalingAction); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoScalingAction{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch_actions.AutoScalingAction",
		[]interface{}{stepScalingAction},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoScalingAction_Override(a AutoScalingAction, stepScalingAction awsautoscaling.StepScalingAction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch_actions.AutoScalingAction",
		[]interface{}{stepScalingAction},
		a,
	)
}

func (a *jsiiProxy_AutoScalingAction) Bind(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) *awscloudwatch.AlarmActionConfig {
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

