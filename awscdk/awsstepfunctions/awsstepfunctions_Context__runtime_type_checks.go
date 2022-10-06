//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func validateContext_NumberAtParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateContext_StringAtParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

