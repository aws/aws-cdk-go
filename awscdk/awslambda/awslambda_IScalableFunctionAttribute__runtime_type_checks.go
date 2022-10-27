//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
)

func (i *jsiiProxy_IScalableFunctionAttribute) validateScaleOnScheduleParameters(id *string, actions *awsapplicationautoscaling.ScalingSchedule) error {
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

func (i *jsiiProxy_IScalableFunctionAttribute) validateScaleOnUtilizationParameters(options *UtilizationScalingOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

