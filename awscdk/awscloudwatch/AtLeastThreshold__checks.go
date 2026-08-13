//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"
)

func validateAtLeastThreshold_CountParameters(count *float64) error {
	if count == nil {
		return fmt.Errorf("parameter count is required, but nil was provided")
	}

	return nil
}

func validateAtLeastThreshold_PercentageParameters(percentage *float64) error {
	if percentage == nil {
		return fmt.Errorf("parameter percentage is required, but nil was provided")
	}

	return nil
}

