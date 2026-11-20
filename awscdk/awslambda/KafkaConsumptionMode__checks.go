//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateKafkaConsumptionMode_OfParameters(mode *string) error {
	if mode == nil {
		return fmt.Errorf("parameter mode is required, but nil was provided")
	}

	return nil
}

