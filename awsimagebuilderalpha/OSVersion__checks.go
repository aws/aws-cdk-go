//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"
)

func validateOSVersion_CustomParameters(platform Platform) error {
	if platform == "" {
		return fmt.Errorf("parameter platform is required, but nil was provided")
	}

	return nil
}

func validateNewOSVersionParameters(platform Platform) error {
	if platform == "" {
		return fmt.Errorf("parameter platform is required, but nil was provided")
	}

	return nil
}

