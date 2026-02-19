//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func validateDefaultEncryptedResourceFactories_GetParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func validateDefaultEncryptedResourceFactories_HasParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func validateDefaultEncryptedResourceFactories_SetParameters(type_ *string, factory IEncryptedResourceFactory) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if factory == nil {
		return fmt.Errorf("parameter factory is required, but nil was provided")
	}

	return nil
}

