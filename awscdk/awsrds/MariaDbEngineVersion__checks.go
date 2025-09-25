//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"
)

func validateMariaDbEngineVersion_OfParameters(mariaDbFullVersion *string, mariaDbMajorVersion *string) error {
	if mariaDbFullVersion == nil {
		return fmt.Errorf("parameter mariaDbFullVersion is required, but nil was provided")
	}

	if mariaDbMajorVersion == nil {
		return fmt.Errorf("parameter mariaDbMajorVersion is required, but nil was provided")
	}

	return nil
}

