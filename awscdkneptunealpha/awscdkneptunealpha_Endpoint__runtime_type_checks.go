//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

import (
	"fmt"
)

func validateNewEndpointParameters(address *string, port *float64) error {
	if address == nil {
		return fmt.Errorf("parameter address is required, but nil was provided")
	}

	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}

	return nil
}

