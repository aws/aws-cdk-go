//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func validateDefaultPolicyFactories_GetParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func validateDefaultPolicyFactories_HasParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func validateDefaultPolicyFactories_SetParameters(type_ *string, factory IResourcePolicyFactory) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if factory == nil {
		return fmt.Errorf("parameter factory is required, but nil was provided")
	}

	return nil
}

