//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"
)

func validateGuardrail_EnableParameters(identifier *string, version *float64) error {
	if identifier == nil {
		return fmt.Errorf("parameter identifier is required, but nil was provided")
	}

	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateGuardrail_EnableDraftParameters(identifier *string) error {
	if identifier == nil {
		return fmt.Errorf("parameter identifier is required, but nil was provided")
	}

	return nil
}

