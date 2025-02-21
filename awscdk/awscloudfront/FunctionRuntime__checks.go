//go:build !no_runtime_type_checking

package awscloudfront

import (
	"fmt"
)

func validateFunctionRuntime_CustomParameters(runtimeString *string) error {
	if runtimeString == nil {
		return fmt.Errorf("parameter runtimeString is required, but nil was provided")
	}

	return nil
}

