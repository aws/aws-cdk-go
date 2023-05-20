//go:build !no_runtime_type_checking

package awscdkneptunealpha

import (
	"fmt"
)

func validateNewEngineVersionParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

