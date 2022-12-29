//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func validateLogFormat_CustomParameters(formatString *string) error {
	if formatString == nil {
		return fmt.Errorf("parameter formatString is required, but nil was provided")
	}

	return nil
}

func validateLogFormat_FieldParameters(field *string) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}

	return nil
}

func validateNewLogFormatParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

