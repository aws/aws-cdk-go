//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"
)

func validateOperatingSystemFamily_OfParameters(family *string) error {
	if family == nil {
		return fmt.Errorf("parameter family is required, but nil was provided")
	}

	return nil
}

