//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsmsk

import (
	"fmt"
)

func validateKafkaVersion_OfParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

