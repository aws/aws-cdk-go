//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"
)

func validateAudioDashRole_OfParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

