//go:build !no_runtime_type_checking

package awsappconfig

import (
	"fmt"
)

func validateJsonSchemaValidator_FromFileParameters(inputPath *string) error {
	if inputPath == nil {
		return fmt.Errorf("parameter inputPath is required, but nil was provided")
	}

	return nil
}

func validateJsonSchemaValidator_FromInlineParameters(code *string) error {
	if code == nil {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	return nil
}

