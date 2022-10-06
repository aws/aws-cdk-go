//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsapprunner

import (
	"fmt"
)

func validateCpu_OfParameters(unit *string) error {
	if unit == nil {
		return fmt.Errorf("parameter unit is required, but nil was provided")
	}

	return nil
}

