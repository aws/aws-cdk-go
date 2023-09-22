//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateSchedule_ProtectedAtParameters(date *time.Time) error {
	if date == nil {
		return fmt.Errorf("parameter date is required, but nil was provided")
	}

	return nil
}

func validateSchedule_ProtectedCronParameters(options *CronOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSchedule_ProtectedExpressionParameters(expression *string) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateSchedule_ProtectedRateParameters(duration Duration) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

