//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"
)

func validateAudioSampleRate_OfParameters(hz *float64) error {
	if hz == nil {
		return fmt.Errorf("parameter hz is required, but nil was provided")
	}

	return nil
}

