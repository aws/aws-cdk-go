//go:build !no_runtime_type_checking

package awsdynamodb

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
)

func (i *jsiiProxy_IScalableTableAttribute) validateScaleOnScheduleParameters(id *string, actions *awsapplicationautoscaling.ScalingSchedule) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if actions == nil {
		return fmt.Errorf("parameter actions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(actions, func() string { return "parameter actions" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IScalableTableAttribute) validateScaleOnUtilizationParameters(props *UtilizationScalingProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

