//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func validateTimeZone_OfParameters(timezoneName *string) error {
	if timezoneName == nil {
		return fmt.Errorf("parameter timezoneName is required, but nil was provided")
	}

	return nil
}

