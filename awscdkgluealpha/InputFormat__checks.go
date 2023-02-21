//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

import (
	"fmt"
)

func validateNewInputFormatParameters(className *string) error {
	if className == nil {
		return fmt.Errorf("parameter className is required, but nil was provided")
	}

	return nil
}

