//go:build !no_runtime_type_checking

package awscloudwatchactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AutoScalingAction) validateBindParameters(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _alarm == nil {
		return fmt.Errorf("parameter _alarm is required, but nil was provided")
	}

	return nil
}

func validateNewAutoScalingActionParameters(stepScalingAction awsautoscaling.StepScalingAction) error {
	if stepScalingAction == nil {
		return fmt.Errorf("parameter stepScalingAction is required, but nil was provided")
	}

	return nil
}

