//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateAcceleratorMemorySize_OfParameters(gibibytes *float64) error {
	if gibibytes == nil {
		return fmt.Errorf("parameter gibibytes is required, but nil was provided")
	}

	return nil
}

