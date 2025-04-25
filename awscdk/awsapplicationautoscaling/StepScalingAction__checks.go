//go:build !no_runtime_type_checking

package awsapplicationautoscaling

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_StepScalingAction) validateAddAdjustmentParameters(adjustment *AdjustmentTier) error {
	if adjustment == nil {
		return fmt.Errorf("parameter adjustment is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(adjustment, func() string { return "parameter adjustment" }); err != nil {
		return err
	}

	return nil
}

func validateStepScalingAction_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewStepScalingActionParameters(scope constructs.Construct, id *string, props *StepScalingActionProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

