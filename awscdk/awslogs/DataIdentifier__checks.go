//go:build !no_runtime_type_checking

package awslogs

import (
	"fmt"
)

func validateNewDataIdentifierParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

