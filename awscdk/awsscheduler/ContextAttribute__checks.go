//go:build !no_runtime_type_checking

package awsscheduler

import (
	"fmt"
)

func validateContextAttribute_FromNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

