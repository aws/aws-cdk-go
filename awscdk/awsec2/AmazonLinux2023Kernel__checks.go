//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func validateNewAmazonLinux2023KernelParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

