//go:build !no_runtime_type_checking

package awseksv2

import (
	"fmt"
)

func validateControlPlaneScalingTier_OfParameters(tier *string) error {
	if tier == nil {
		return fmt.Errorf("parameter tier is required, but nil was provided")
	}

	return nil
}

