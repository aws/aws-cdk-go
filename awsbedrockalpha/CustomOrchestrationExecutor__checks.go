//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func validateCustomOrchestrationExecutor_FromLambdaParameters(lambdaFunction awslambda.IFunction) error {
	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}

	return nil
}

