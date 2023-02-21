//go:build !no_runtime_type_checking

// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
)

func validateNewSetVariableActionParameters(variableName *string, value awscdkioteventsalpha.Expression) error {
	if variableName == nil {
		return fmt.Errorf("parameter variableName is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

