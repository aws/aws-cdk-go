//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"
)

func validateRuntime_OfParameters(runtime *string) error {
	if runtime == nil {
		return fmt.Errorf("parameter runtime is required, but nil was provided")
	}

	return nil
}

