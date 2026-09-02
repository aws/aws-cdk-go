//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"
)

func validateCaptionFontSize_OfParameters(points *float64) error {
	if points == nil {
		return fmt.Errorf("parameter points is required, but nil was provided")
	}

	return nil
}

