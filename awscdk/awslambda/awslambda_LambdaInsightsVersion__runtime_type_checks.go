//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateLambdaInsightsVersion_FromInsightVersionArnParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

