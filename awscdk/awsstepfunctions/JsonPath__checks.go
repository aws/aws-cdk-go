//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func validateJsonPath_ArrayContainsParameters(array interface{}, value interface{}) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_ArrayGetItemParameters(array interface{}, index *float64) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_ArrayLengthParameters(array interface{}) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_ArrayPartitionParameters(array interface{}, chunkSize *float64) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	if chunkSize == nil {
		return fmt.Errorf("parameter chunkSize is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_ArrayRangeParameters(start *float64, end *float64, step *float64) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
	}

	if step == nil {
		return fmt.Errorf("parameter step is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_ArrayUniqueParameters(array interface{}) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_Base64DecodeParameters(base64 *string) error {
	if base64 == nil {
		return fmt.Errorf("parameter base64 is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_Base64EncodeParameters(input *string) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_FormatParameters(formatString *string) error {
	if formatString == nil {
		return fmt.Errorf("parameter formatString is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_HashParameters(data interface{}, algorithm *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	if algorithm == nil {
		return fmt.Errorf("parameter algorithm is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_IsEncodedJsonPathParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_JsonMergeParameters(value1 interface{}, value2 interface{}) error {
	if value1 == nil {
		return fmt.Errorf("parameter value1 is required, but nil was provided")
	}

	if value2 == nil {
		return fmt.Errorf("parameter value2 is required, but nil was provided")
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

func validateJsonPath_MathAddParameters(num1 *float64, num2 *float64) error {
	if num1 == nil {
		return fmt.Errorf("parameter num1 is required, but nil was provided")
	}

	if num2 == nil {
		return fmt.Errorf("parameter num2 is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_MathRandomParameters(start *float64, end *float64) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
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

func validateJsonPath_StringSplitParameters(inputString *string, splitter *string) error {
	if inputString == nil {
		return fmt.Errorf("parameter inputString is required, but nil was provided")
	}

	if splitter == nil {
		return fmt.Errorf("parameter splitter is required, but nil was provided")
	}

	return nil
}

func validateJsonPath_StringToJsonParameters(jsonString *string) error {
	if jsonString == nil {
		return fmt.Errorf("parameter jsonString is required, but nil was provided")
	}

	return nil
}

