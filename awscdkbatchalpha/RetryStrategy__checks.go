//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"fmt"
)

func validateRetryStrategy_OfParameters(action Action, on Reason) error {
	if action == "" {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	if on == nil {
		return fmt.Errorf("parameter on is required, but nil was provided")
	}

	return nil
}

func validateNewRetryStrategyParameters(action Action, on Reason) error {
	if action == "" {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	if on == nil {
		return fmt.Errorf("parameter on is required, but nil was provided")
	}

	return nil
}

