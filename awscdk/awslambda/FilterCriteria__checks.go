//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateFilterCriteria_FilterParameters(filter *map[string]interface{}) error {
	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}

	return nil
}

