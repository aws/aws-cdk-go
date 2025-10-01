//go:build !no_runtime_type_checking

package awssigner

import (
	"fmt"
)

func validatePlatform_OfParameters(platformId *string) error {
	if platformId == nil {
		return fmt.Errorf("parameter platformId is required, but nil was provided")
	}

	return nil
}

