//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

import (
	"fmt"
)

func validateNewParameterGroupFamilyParameters(family *string) error {
	if family == nil {
		return fmt.Errorf("parameter family is required, but nil was provided")
	}

	return nil
}

