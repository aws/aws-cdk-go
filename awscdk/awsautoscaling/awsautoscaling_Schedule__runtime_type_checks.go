//go:build !no_runtime_type_checking

package awsautoscaling

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateSchedule_CronParameters(options *CronOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSchedule_ExpressionParameters(expression *string) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

