//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"
)

func validateGopSize_FramesParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateGopSize_SecondsParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

