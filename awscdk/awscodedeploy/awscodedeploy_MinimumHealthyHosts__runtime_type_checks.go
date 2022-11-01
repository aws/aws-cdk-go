//go:build !no_runtime_type_checking

package awscodedeploy

import (
	"fmt"
)

func validateMinimumHealthyHosts_CountParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMinimumHealthyHosts_PercentageParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

