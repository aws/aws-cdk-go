//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

func validateAcceleratorType_OfParameters(acceleratorClass AcceleratorClass, instanceSize awsec2.InstanceSize) error {
	if acceleratorClass == nil {
		return fmt.Errorf("parameter acceleratorClass is required, but nil was provided")
	}

	if instanceSize == "" {
		return fmt.Errorf("parameter instanceSize is required, but nil was provided")
	}

	return nil
}

func validateNewAcceleratorTypeParameters(instanceTypeIdentifier *string) error {
	if instanceTypeIdentifier == nil {
		return fmt.Errorf("parameter instanceTypeIdentifier is required, but nil was provided")
	}

	return nil
}

