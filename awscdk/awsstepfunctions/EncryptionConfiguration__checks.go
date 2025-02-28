//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func (j *jsiiProxy_EncryptionConfiguration) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEncryptionConfigurationParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

