//go:build !no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	"fmt"
)

func validateExpectedResult_ArrayWithParameters(expected *[]interface{}) error {
	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func validateExpectedResult_ExactParameters(expected interface{}) error {
	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func validateExpectedResult_ObjectLikeParameters(expected *map[string]interface{}) error {
	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func validateExpectedResult_StringLikeRegexpParameters(expected *string) error {
	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ExpectedResult) validateSetResultParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

