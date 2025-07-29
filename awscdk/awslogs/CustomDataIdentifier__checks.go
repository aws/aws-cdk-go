//go:build !no_runtime_type_checking

package awslogs

import (
	"fmt"
)

func validateNewCustomDataIdentifierParameters(name *string, regex *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if regex == nil {
		return fmt.Errorf("parameter regex is required, but nil was provided")
	}

	return nil
}

