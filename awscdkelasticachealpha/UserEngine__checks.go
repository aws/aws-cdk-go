//go:build !no_runtime_type_checking

package awscdkelasticachealpha

import (
	"fmt"
)

func validateUserEngine_OfParameters(engineType *string) error {
	if engineType == nil {
		return fmt.Errorf("parameter engineType is required, but nil was provided")
	}

	return nil
}

