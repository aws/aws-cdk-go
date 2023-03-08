//go:build !no_runtime_type_checking

package awss3

import (
	"fmt"
)

func validateReplaceKey_PrefixWithParameters(keyReplacement *string) error {
	if keyReplacement == nil {
		return fmt.Errorf("parameter keyReplacement is required, but nil was provided")
	}

	return nil
}

func validateReplaceKey_WithParameters(keyReplacement *string) error {
	if keyReplacement == nil {
		return fmt.Errorf("parameter keyReplacement is required, but nil was provided")
	}

	return nil
}

