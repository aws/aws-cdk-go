//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"
)

func validateCaCertificate_OfParameters(identifier *string) error {
	if identifier == nil {
		return fmt.Errorf("parameter identifier is required, but nil was provided")
	}

	return nil
}

