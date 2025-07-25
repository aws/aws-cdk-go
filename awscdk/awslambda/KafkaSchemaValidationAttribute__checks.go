//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateKafkaSchemaValidationAttribute_OfParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

