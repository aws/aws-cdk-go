//go:build !no_runtime_type_checking

package awscdkbatchalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func (i *jsiiProxy_IManagedEc2EcsComputeEnvironment) validateAddInstanceClassParameters(instanceClass awsec2.InstanceClass) error {
	if instanceClass == "" {
		return fmt.Errorf("parameter instanceClass is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IManagedEc2EcsComputeEnvironment) validateAddInstanceTypeParameters(instanceType awsec2.InstanceType) error {
	if instanceType == nil {
		return fmt.Errorf("parameter instanceType is required, but nil was provided")
	}

	return nil
}

