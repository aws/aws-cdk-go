//go:build !no_runtime_type_checking

package awscdkapplicationsignalsalpha

import (
	"fmt"
)

func validateNewPythonInstrumentationVersionParameters(imageRepo *string, version *string, memoryLimit *float64) error {
	if imageRepo == nil {
		return fmt.Errorf("parameter imageRepo is required, but nil was provided")
	}

	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	if memoryLimit == nil {
		return fmt.Errorf("parameter memoryLimit is required, but nil was provided")
	}

	return nil
}

