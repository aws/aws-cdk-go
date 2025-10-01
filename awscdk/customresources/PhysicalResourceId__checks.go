//go:build !no_runtime_type_checking

package customresources

import (
	"fmt"
)

func validatePhysicalResourceId_FromResponseParameters(responsePath *string) error {
	if responsePath == nil {
		return fmt.Errorf("parameter responsePath is required, but nil was provided")
	}

	return nil
}

func validatePhysicalResourceId_OfParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

