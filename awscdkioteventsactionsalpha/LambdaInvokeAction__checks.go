//go:build !no_runtime_type_checking

// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func validateNewLambdaInvokeActionParameters(func_ awslambda.IFunction) error {
	if func_ == nil {
		return fmt.Errorf("parameter func_ is required, but nil was provided")
	}

	return nil
}

