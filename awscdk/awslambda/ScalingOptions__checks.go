//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateScalingOptions_ManualParameters(scalingPolicies *[]TargetTrackingScalingPolicy) error {
	if scalingPolicies == nil {
		return fmt.Errorf("parameter scalingPolicies is required, but nil was provided")
	}

	return nil
}

