//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"
)

func (v *jsiiProxy_ValidationResults) validateCollectParameters(result ValidationResult) error {
	if result == nil {
		return fmt.Errorf("parameter result is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_ValidationResults) validateWrapParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ValidationResults) validateSetResultsParameters(val *[]ValidationResult) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

