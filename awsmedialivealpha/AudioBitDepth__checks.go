//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"
)

func validateAudioBitDepth_OfParameters(bits *float64) error {
	if bits == nil {
		return fmt.Errorf("parameter bits is required, but nil was provided")
	}

	return nil
}

