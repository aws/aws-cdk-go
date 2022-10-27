//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (s *jsiiProxy_StackOutputReference) validateIsProducedByParameters(stack StackDeployment) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

func validateStackOutputReference_FromCfnOutputParameters(output awscdk.CfnOutput) error {
	if output == nil {
		return fmt.Errorf("parameter output is required, but nil was provided")
	}

	return nil
}

