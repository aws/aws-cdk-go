//go:build !no_runtime_type_checking

package awscodepipelineactions

import (
	"fmt"
)

func validateEc2MaxInstances_PercentageParameters(percentage *float64) error {
	if percentage == nil {
		return fmt.Errorf("parameter percentage is required, but nil was provided")
	}

	return nil
}

func validateEc2MaxInstances_TargetsParameters(targets *float64) error {
	if targets == nil {
		return fmt.Errorf("parameter targets is required, but nil was provided")
	}

	return nil
}

