//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func validateDependableTrait_GetParameters(instance IDependable) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	return nil
}

func validateDependableTrait_ImplementParameters(instance IDependable, trait DependableTrait) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	if trait == nil {
		return fmt.Errorf("parameter trait is required, but nil was provided")
	}

	return nil
}

