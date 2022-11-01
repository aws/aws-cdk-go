//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func validateNewInitConfigParameters(elements *[]InitElement) error {
	if elements == nil {
		return fmt.Errorf("parameter elements is required, but nil was provided")
	}

	return nil
}

