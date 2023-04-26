//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func validateJsonPath_FormatParameters(formatString *string) error {
	if formatString == nil {
		return fmt.Errorf("parameter formatString is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_IsEncodedJsonPathParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_JsonToStringParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_ListAtParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_NumberAtParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_ObjectAtParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_StringAtParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_StringToJsonParameters(jsonString *string) error {
	if jsonString == nil {
		return fmt.Errorf("parameter jsonString is required, but nil was provided")
	}

	return nil
}

