//go:build !no_runtime_type_checking

package awsbatch

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch"
)

func (i *jsiiProxy_IJobQueue) validateAddComputeEnvironmentParameters(computeEnvironment interfacesawsbatch.IComputeEnvironmentRef, order *float64) error {
	if computeEnvironment == nil {
		return fmt.Errorf("parameter computeEnvironment is required, but nil was provided")
	}

	if order == nil {
		return fmt.Errorf("parameter order is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IJobQueue) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

