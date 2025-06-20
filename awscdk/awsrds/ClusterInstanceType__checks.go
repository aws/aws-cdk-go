//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"
)

func validateNewClusterInstanceTypeParameters(instanceType *string, type_ InstanceType) error {
	if instanceType == nil {
		return fmt.Errorf("parameter instanceType is required, but nil was provided")
	}

	if type_ == "" {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

