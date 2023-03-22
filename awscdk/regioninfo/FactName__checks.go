//go:build !no_runtime_type_checking

package regioninfo

import (
	"fmt"
)

func validateFactName_AdotLambdaLayerParameters(type_ *string, version *string, architecture *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	if architecture == nil {
		return fmt.Errorf("parameter architecture is required, but nil was provided")
	}

	return nil
}

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

