//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (v *jsiiProxy_ValidationResult) validatePrefixParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

