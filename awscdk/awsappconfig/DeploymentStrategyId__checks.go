//go:build !no_runtime_type_checking

package awsappconfig

import (
	"fmt"
)

func validateDeploymentStrategyId_FromStringParameters(deploymentStrategyId *string) error {
	if deploymentStrategyId == nil {
		return fmt.Errorf("parameter deploymentStrategyId is required, but nil was provided")
	}

	return nil
}

