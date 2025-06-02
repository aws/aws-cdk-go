//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func validateComparablePrincipal_DedupeStringForParameters(x IPrincipal) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateComparablePrincipal_IsComparablePrincipalParameters(x IPrincipal) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

