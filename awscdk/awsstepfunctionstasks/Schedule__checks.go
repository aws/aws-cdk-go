//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"
	"time"

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

func validateSchedule_OneTimeParameters(time *time.Time) error {
	if time == nil {
		return fmt.Errorf("parameter time is required, but nil was provided")
	}

	return nil
}

func validateSchedule_RateParameters(duration awscdk.Duration) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

