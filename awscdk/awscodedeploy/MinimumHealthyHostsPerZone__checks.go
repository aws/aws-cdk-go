//go:build !no_runtime_type_checking

package awscodedeploy

import (
	"fmt"
)

func validateMinimumHealthyHostsPerZone_CountParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMinimumHealthyHostsPerZone_PercentageParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

