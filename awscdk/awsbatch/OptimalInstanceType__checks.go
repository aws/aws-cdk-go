//go:build !no_runtime_type_checking

package awsbatch

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func (o *jsiiProxy_OptimalInstanceType) validateSameInstanceClassAsParameters(other awsec2.InstanceType) error {
	if other == nil {
		return fmt.Errorf("parameter other is required, but nil was provided")
	}

	return nil
}

func validateOptimalInstanceType_OfParameters(instanceClass awsec2.InstanceClass, instanceSize awsec2.InstanceSize) error {
	if instanceClass == "" {
		return fmt.Errorf("parameter instanceClass is required, but nil was provided")
	}

	if instanceSize == "" {
		return fmt.Errorf("parameter instanceSize is required, but nil was provided")
	}

	return nil
}

