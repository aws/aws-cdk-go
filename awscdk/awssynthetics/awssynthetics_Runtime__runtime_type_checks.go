//go:build !no_runtime_type_checking

package awssynthetics

import (
	"fmt"
)

func validateNewRuntimeParameters(name *string, family RuntimeFamily) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if family == "" {
		return fmt.Errorf("parameter family is required, but nil was provided")
	}

	return nil
}

