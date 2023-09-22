//go:build !no_runtime_type_checking

package awscdkscheduleralpha

import (
	"fmt"
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func validateScheduleExpression_AtParameters(date *time.Time) error {
	if date == nil {
		return fmt.Errorf("parameter date is required, but nil was provided")
	}

	return nil
}

func validateScheduleExpression_CronParameters(options *awscdk.CronOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateScheduleExpression_ExpressionParameters(expression *string) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateScheduleExpression_ProtectedAtParameters(date *time.Time) error {
	if date == nil {
		return fmt.Errorf("parameter date is required, but nil was provided")
	}

	return nil
}

func validateScheduleExpression_ProtectedCronParameters(options *awscdk.CronOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateScheduleExpression_ProtectedExpressionParameters(expression *string) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateScheduleExpression_ProtectedRateParameters(duration awscdk.Duration) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

func validateScheduleExpression_RateParameters(duration awscdk.Duration) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

