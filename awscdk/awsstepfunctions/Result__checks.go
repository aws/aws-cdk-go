//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func validateResult_FromArrayParameters(value *[]interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateResult_FromBooleanParameters(value *bool) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateResult_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateResult_FromObjectParameters(value *map[string]interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateResult_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateNewResultParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

