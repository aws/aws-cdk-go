//go:build !no_runtime_type_checking

package customresources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CustomResourceLambdaRuntime) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func validateNewCustomResourceLambdaRuntimeParameters(lambdaRuntime awslambda.Runtime) error {
	if lambdaRuntime == nil {
		return fmt.Errorf("parameter lambdaRuntime is required, but nil was provided")
	}

	return nil
}

