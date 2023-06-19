//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func validateNewScopedAwsParameters(scope Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

