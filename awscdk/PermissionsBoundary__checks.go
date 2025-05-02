//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func validatePermissionsBoundary_FromArnParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func validatePermissionsBoundary_FromNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

