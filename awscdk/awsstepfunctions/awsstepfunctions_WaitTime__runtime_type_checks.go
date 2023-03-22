//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func validateWaitTime_DurationParameters(duration awscdk.Duration) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

func validateWaitTime_SecondsPathParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateWaitTime_TimestampParameters(timestamp *string) error {
	if timestamp == nil {
		return fmt.Errorf("parameter timestamp is required, but nil was provided")
	}

	return nil
}

func validateWaitTime_TimestampPathParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

