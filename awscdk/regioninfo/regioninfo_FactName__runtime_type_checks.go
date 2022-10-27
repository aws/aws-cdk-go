//go:build !no_runtime_type_checking

package regioninfo

import (
	"fmt"
)

func validateFactName_CloudwatchLambdaInsightsVersionParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateFactName_ServicePrincipalParameters(service *string) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

