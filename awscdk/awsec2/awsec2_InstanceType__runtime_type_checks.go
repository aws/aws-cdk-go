//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func validateInstanceType_OfParameters(instanceClass InstanceClass, instanceSize InstanceSize) error {
	if instanceClass == "" {
		return fmt.Errorf("parameter instanceClass is required, but nil was provided")
	}

	if instanceSize == "" {
		return fmt.Errorf("parameter instanceSize is required, but nil was provided")
	}

	return nil
}

func validateNewInstanceTypeParameters(instanceTypeIdentifier *string) error {
	if instanceTypeIdentifier == nil {
		return fmt.Errorf("parameter instanceTypeIdentifier is required, but nil was provided")
	}

	return nil
}

