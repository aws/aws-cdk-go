//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (g *jsiiProxy_GenericSSMParameterImage) validateGetImageParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateNewGenericSSMParameterImageParameters(parameterName *string, os OperatingSystemType) error {
	if parameterName == nil {
		return fmt.Errorf("parameter parameterName is required, but nil was provided")
	}

	if os == "" {
		return fmt.Errorf("parameter os is required, but nil was provided")
	}

	return nil
}

