//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"fmt"
)

func (i *jsiiProxy_IJobDefinition) validateAddRetryStrategyParameters(strategy RetryStrategy) error {
	if strategy == nil {
		return fmt.Errorf("parameter strategy is required, but nil was provided")
	}

	return nil
}

