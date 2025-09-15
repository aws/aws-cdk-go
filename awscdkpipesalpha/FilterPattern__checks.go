//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"
)

func validateFilterPattern_FromObjectParameters(patternObject *map[string]interface{}) error {
	if patternObject == nil {
		return fmt.Errorf("parameter patternObject is required, but nil was provided")
	}

	return nil
}

