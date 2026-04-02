//go:build !no_runtime_type_checking

package awsmediapackagev2alpha

import (
	"fmt"
)

func validateHeadersCMSD_OfParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

