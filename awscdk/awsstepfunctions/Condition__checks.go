//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func validateCondition_BooleanEqualsParameters(variable *string, value *bool) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_BooleanEqualsJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsBooleanParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsNotBooleanParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsNotNullParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsNotNumericParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsNotPresentParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsNotStringParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsNotTimestampParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsNullParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsNumericParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsPresentParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsStringParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_IsTimestampParameters(variable *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	return nil
}

func validateCondition_NotParameters(condition Condition) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	return nil
}

func validateCondition_NumberEqualsParameters(variable *string, value *float64) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_NumberEqualsJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_NumberGreaterThanParameters(variable *string, value *float64) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_NumberGreaterThanEqualsParameters(variable *string, value *float64) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_NumberGreaterThanEqualsJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_NumberGreaterThanJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_NumberLessThanParameters(variable *string, value *float64) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_NumberLessThanEqualsParameters(variable *string, value *float64) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_NumberLessThanEqualsJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_NumberLessThanJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringEqualsParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringEqualsJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringGreaterThanParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringGreaterThanEqualsParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringGreaterThanEqualsJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringGreaterThanJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringLessThanParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringLessThanEqualsParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringLessThanEqualsJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringLessThanJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_StringMatchesParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_TimestampEqualsParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_TimestampEqualsJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_TimestampGreaterThanParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_TimestampGreaterThanEqualsParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_TimestampGreaterThanEqualsJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_TimestampGreaterThanJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_TimestampLessThanParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_TimestampLessThanEqualsParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_TimestampLessThanEqualsJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCondition_TimestampLessThanJsonPathParameters(variable *string, value *string) error {
	if variable == nil {
		return fmt.Errorf("parameter variable is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

