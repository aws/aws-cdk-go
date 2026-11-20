//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateTargetTrackingScalingPolicy_CpuUtilizationParameters(targetCpuUtilization *float64) error {
	if targetCpuUtilization == nil {
		return fmt.Errorf("parameter targetCpuUtilization is required, but nil was provided")
	}

	return nil
}

func validateTargetTrackingScalingPolicy_GpuUtilizationParameters(targetGpuUtilization *float64) error {
	if targetGpuUtilization == nil {
		return fmt.Errorf("parameter targetGpuUtilization is required, but nil was provided")
	}

	return nil
}

