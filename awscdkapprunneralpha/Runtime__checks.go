//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	"fmt"
)

func validateRuntime_OfParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

