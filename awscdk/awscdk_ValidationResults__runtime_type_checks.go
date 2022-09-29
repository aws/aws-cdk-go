//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
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

