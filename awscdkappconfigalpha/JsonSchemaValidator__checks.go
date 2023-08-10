//go:build !no_runtime_type_checking

package awscdkappconfigalpha

import (
	"fmt"
)

func validateJsonSchemaValidator_FromFileParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateJsonSchemaValidator_FromInlineParameters(code *string) error {
	if code == nil {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	return nil
}

