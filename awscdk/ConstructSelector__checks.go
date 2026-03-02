//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func validateConstructSelector_ByIdParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateConstructSelector_ByPathParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

