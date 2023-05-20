//go:build !no_runtime_type_checking

package awscdkbatchalpha

import (
	"fmt"
)

func (i *jsiiProxy_IJobQueue) validateAddComputeEnvironmentParameters(computeEnvironment IComputeEnvironment, order *float64) error {
	if computeEnvironment == nil {
		return fmt.Errorf("parameter computeEnvironment is required, but nil was provided")
	}

	if order == nil {
		return fmt.Errorf("parameter order is required, but nil was provided")
	}

	return nil
}

