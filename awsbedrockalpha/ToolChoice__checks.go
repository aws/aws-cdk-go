//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"
)

func validateToolChoice_SpecificToolParameters(toolName *string) error {
	if toolName == nil {
		return fmt.Errorf("parameter toolName is required, but nil was provided")
	}

	return nil
}

func validateNewToolChoiceParameters(any interface{}, auto interface{}) error {
	if any == nil {
		return fmt.Errorf("parameter any is required, but nil was provided")
	}

	if auto == nil {
		return fmt.Errorf("parameter auto is required, but nil was provided")
	}

	return nil
}

