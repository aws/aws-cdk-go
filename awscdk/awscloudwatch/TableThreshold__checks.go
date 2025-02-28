//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"
)

func validateTableThreshold_AboveParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateTableThreshold_BelowParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateTableThreshold_BetweenParameters(lowerBound *float64, upperBound *float64) error {
	if lowerBound == nil {
		return fmt.Errorf("parameter lowerBound is required, but nil was provided")
	}

	if upperBound == nil {
		return fmt.Errorf("parameter upperBound is required, but nil was provided")
	}

	return nil
}

