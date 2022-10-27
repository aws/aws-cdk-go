//go:build !no_runtime_type_checking

package awsapprunner

import (
	"fmt"
)

func validateMemory_OfParameters(unit *string) error {
	if unit == nil {
		return fmt.Errorf("parameter unit is required, but nil was provided")
	}

	return nil
}

