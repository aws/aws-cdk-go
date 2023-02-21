//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::MSK
package awscdkmskalpha

import (
	"fmt"
)

func validateKafkaVersion_OfParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

