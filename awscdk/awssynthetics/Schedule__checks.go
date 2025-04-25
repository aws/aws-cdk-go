//go:build !no_runtime_type_checking

package awssynthetics

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
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

func validateSchedule_RateParameters(interval awscdk.Duration) error {
	if interval == nil {
		return fmt.Errorf("parameter interval is required, but nil was provided")
	}

	return nil
}

