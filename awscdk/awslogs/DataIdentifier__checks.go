//go:build !no_runtime_type_checking

package awslogs

import (
	"fmt"
)

func validateNewDataIdentifierParameters(identifier *string) error {
	if identifier == nil {
		return fmt.Errorf("parameter identifier is required, but nil was provided")
	}

	return nil
}

