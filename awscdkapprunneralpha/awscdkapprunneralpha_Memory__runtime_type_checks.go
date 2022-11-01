//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	"fmt"
)

func validateMemory_OfParameters(unit *string) error {
	if unit == nil {
		return fmt.Errorf("parameter unit is required, but nil was provided")
	}

	return nil
}

