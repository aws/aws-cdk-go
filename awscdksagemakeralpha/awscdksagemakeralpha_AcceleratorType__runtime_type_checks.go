//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

import (
	"fmt"
)

func validateAcceleratorType_OfParameters(acceleratorType *string) error {
	if acceleratorType == nil {
		return fmt.Errorf("parameter acceleratorType is required, but nil was provided")
	}

	return nil
}

func validateNewAcceleratorTypeParameters(acceleratorType *string) error {
	if acceleratorType == nil {
		return fmt.Errorf("parameter acceleratorType is required, but nil was provided")
	}

	return nil
}

