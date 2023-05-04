//go:build !no_runtime_type_checking

package awscognito

import (
	"fmt"
)

func validateProviderAttribute_OtherParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

