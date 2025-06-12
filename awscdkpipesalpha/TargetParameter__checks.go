//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"
)

func validateTargetParameter_FromJsonPathParameters(jsonPath *string) error {
	if jsonPath == nil {
		return fmt.Errorf("parameter jsonPath is required, but nil was provided")
	}

	return nil
}

