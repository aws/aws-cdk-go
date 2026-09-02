//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"
)

func validateFramerate_OfParameters(numerator *float64, denominator *float64) error {
	if numerator == nil {
		return fmt.Errorf("parameter numerator is required, but nil was provided")
	}

	if denominator == nil {
		return fmt.Errorf("parameter denominator is required, but nil was provided")
	}

	return nil
}

