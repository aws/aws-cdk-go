//go:build !no_runtime_type_checking

package awsecrassets

import (
	"fmt"
)

func validateNetworkMode_CustomParameters(mode *string) error {
	if mode == nil {
		return fmt.Errorf("parameter mode is required, but nil was provided")
	}

	return nil
}

func validateNetworkMode_FromContainerParameters(containerId *string) error {
	if containerId == nil {
		return fmt.Errorf("parameter containerId is required, but nil was provided")
	}

	return nil
}

