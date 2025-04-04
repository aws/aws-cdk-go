//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func validateTaskInput_FromJsonPathAtParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateTaskInput_FromObjectParameters(obj *map[string]interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func validateTaskInput_FromTextParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

