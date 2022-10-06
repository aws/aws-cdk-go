//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsapigatewayv2

import (
	"fmt"
)

func validateMappingValue_ContextVariableParameters(variableName *string) error {
	if variableName == nil {
		return fmt.Errorf("parameter variableName is required, but nil was provided")
	}

	return nil
}

func validateMappingValue_CustomParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMappingValue_RequestBodyParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateMappingValue_RequestHeaderParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateMappingValue_RequestPathParamParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateMappingValue_RequestQueryStringParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateMappingValue_StageVariableParameters(variableName *string) error {
	if variableName == nil {
		return fmt.Errorf("parameter variableName is required, but nil was provided")
	}

	return nil
}

func validateNewMappingValueParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

