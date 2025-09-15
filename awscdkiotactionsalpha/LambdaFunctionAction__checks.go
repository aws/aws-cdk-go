//go:build !no_runtime_type_checking

package awscdkiotactionsalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func validateNewLambdaFunctionActionParameters(func_ awslambda.IFunction) error {
	if func_ == nil {
		return fmt.Errorf("parameter func_ is required, but nil was provided")
	}

	return nil
}

