//go:build !no_runtime_type_checking

package awscdksagemakeralpha

import (
	"fmt"
)

func validateInstanceType_OfParameters(instanceType *string) error {
	if instanceType == nil {
		return fmt.Errorf("parameter instanceType is required, but nil was provided")
	}

	return nil
}

func validateNewInstanceTypeParameters(instanceType *string) error {
	if instanceType == nil {
		return fmt.Errorf("parameter instanceType is required, but nil was provided")
	}

	return nil
}

