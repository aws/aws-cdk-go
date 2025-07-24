//go:build !no_runtime_type_checking

package awscdkiotalpha

import (
	"fmt"
)

func validateDayOfMonth_OfParameters(day *float64) error {
	if day == nil {
		return fmt.Errorf("parameter day is required, but nil was provided")
	}

	return nil
}

