//go:build !no_runtime_type_checking

package cxapi

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEnvironmentPlaceholders_ReplaceParameters(object interface{}, values *EnvironmentPlaceholderValues) error {
	if object == nil {
		return fmt.Errorf("parameter object is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(values, func() string { return "parameter values" }); err != nil {
		return err
	}

	return nil
}

func validateEnvironmentPlaceholders_ReplaceAsyncParameters(object interface{}, provider IEnvironmentPlaceholderProvider) error {
	if object == nil {
		return fmt.Errorf("parameter object is required, but nil was provided")
	}

	if provider == nil {
		return fmt.Errorf("parameter provider is required, but nil was provided")
	}

	return nil
}

