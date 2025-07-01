//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"
)

func validateStorageParameter_ColumnCountMismatchHandlingParameters(value ColumnCountMismatchHandlingAction) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_CompressionTypeParameters(value CompressionType) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_CustomParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_DataCleansingEnabledParameters(value *bool) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_InvalidCharHandlingParameters(value InvalidCharHandlingAction) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_NumericOverflowHandlingParameters(value NumericOverflowHandlingAction) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_NumRowsParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_OrcSchemaResolutionParameters(value OrcColumnMappingType) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_ReplacementCharParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_SerializationNullFormatParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_SkipHeaderLineCountParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_SurplusBytesHandlingParameters(value SurplusBytesHandlingAction) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_SurplusCharHandlingParameters(value SurplusCharHandlingAction) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_WriteKmsKeyIdParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_WriteMaxFileSizeMbParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateStorageParameter_WriteParallelParameters(value WriteParallel) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateNewStorageParameterParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

