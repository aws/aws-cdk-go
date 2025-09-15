//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateJsonPath_ArrayContainsParameters(array interface{}, value interface{}) error {
	return nil
}

func validateJsonPath_ArrayGetItemParameters(array interface{}, index *float64) error {
	return nil
}

func validateJsonPath_ArrayLengthParameters(array interface{}) error {
	return nil
}

func validateJsonPath_ArrayPartitionParameters(array interface{}, chunkSize *float64) error {
	return nil
}

func validateJsonPath_ArrayRangeParameters(start *float64, end *float64, step *float64) error {
	return nil
}

func validateJsonPath_ArrayUniqueParameters(array interface{}) error {
	return nil
}

func validateJsonPath_Base64DecodeParameters(base64 *string) error {
	return nil
}

func validateJsonPath_Base64EncodeParameters(input *string) error {
	return nil
}

func validateJsonPath_FormatParameters(formatString *string) error {
	return nil
}

func validateJsonPath_HashParameters(data interface{}, algorithm *string) error {
	return nil
}

func validateJsonPath_IsEncodedJsonPathParameters(value *string) error {
	return nil
}

func validateJsonPath_JsonMergeParameters(value1 interface{}, value2 interface{}) error {
	return nil
}

func validateJsonPath_JsonToStringParameters(value interface{}) error {
	return nil
}

func validateJsonPath_ListAtParameters(path *string) error {
	return nil
}

func validateJsonPath_MathAddParameters(num1 *float64, num2 *float64) error {
	return nil
}

func validateJsonPath_MathRandomParameters(start *float64, end *float64) error {
	return nil
}

func validateJsonPath_NumberAtParameters(path *string) error {
	return nil
}

func validateJsonPath_ObjectAtParameters(path *string) error {
	return nil
}

func validateJsonPath_StringAtParameters(path *string) error {
	return nil
}

func validateJsonPath_StringSplitParameters(inputString *string, splitter *string) error {
	return nil
}

func validateJsonPath_StringToJsonParameters(jsonString *string) error {
	return nil
}

