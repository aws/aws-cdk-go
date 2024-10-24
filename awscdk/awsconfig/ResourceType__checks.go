//go:build !no_runtime_type_checking

package awsconfig

import (
	"fmt"
)

func validateResourceType_OfParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

