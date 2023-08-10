//go:build !no_runtime_type_checking

package awscdkappconfigalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func validateLambdaValidator_FromFunctionParameters(func_ awslambda.Function) error {
	if func_ == nil {
		return fmt.Errorf("parameter func_ is required, but nil was provided")
	}

	return nil
}

