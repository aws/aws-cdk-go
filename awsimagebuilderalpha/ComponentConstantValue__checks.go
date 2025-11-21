//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"
)

func validateComponentConstantValue_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateNewComponentConstantValueParameters(type_ *string, value interface{}) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

