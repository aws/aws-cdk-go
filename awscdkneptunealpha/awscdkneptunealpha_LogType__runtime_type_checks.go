//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

import (
	"fmt"
)

func validateNewLogTypeParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

