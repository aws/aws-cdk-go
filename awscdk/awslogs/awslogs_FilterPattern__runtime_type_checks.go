//go:build !no_runtime_type_checking

package awslogs

import (
	"fmt"
)

func validateFilterPattern_BooleanValueParameters(jsonField *string, value *bool) error {
	if jsonField == nil {
		return fmt.Errorf("parameter jsonField is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFilterPattern_ExistsParameters(jsonField *string) error {
	if jsonField == nil {
		return fmt.Errorf("parameter jsonField is required, but nil was provided")
	}

	return nil
}

func validateFilterPattern_IsNullParameters(jsonField *string) error {
	if jsonField == nil {
		return fmt.Errorf("parameter jsonField is required, but nil was provided")
	}

	return nil
}

func validateFilterPattern_LiteralParameters(logPatternString *string) error {
	if logPatternString == nil {
		return fmt.Errorf("parameter logPatternString is required, but nil was provided")
	}

	return nil
}

func validateFilterPattern_NotExistsParameters(jsonField *string) error {
	if jsonField == nil {
		return fmt.Errorf("parameter jsonField is required, but nil was provided")
	}

	return nil
}

func validateFilterPattern_NumberValueParameters(jsonField *string, comparison *string, value *float64) error {
	if jsonField == nil {
		return fmt.Errorf("parameter jsonField is required, but nil was provided")
	}

	if comparison == nil {
		return fmt.Errorf("parameter comparison is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFilterPattern_StringValueParameters(jsonField *string, comparison *string, value *string) error {
	if jsonField == nil {
		return fmt.Errorf("parameter jsonField is required, but nil was provided")
	}

	if comparison == nil {
		return fmt.Errorf("parameter comparison is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

