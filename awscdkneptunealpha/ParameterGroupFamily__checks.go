//go:build !no_runtime_type_checking

package awscdkneptunealpha

import (
	"fmt"
)

func validateNewParameterGroupFamilyParameters(family *string) error {
	if family == nil {
		return fmt.Errorf("parameter family is required, but nil was provided")
	}

	return nil
}

