//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"
)

func validateDefaultValue_ValueParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

