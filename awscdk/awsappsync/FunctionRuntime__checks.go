//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func validateNewFunctionRuntimeParameters(family FunctionRuntimeFamily, version *string) error {
	if family == "" {
		return fmt.Errorf("parameter family is required, but nil was provided")
	}

	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

