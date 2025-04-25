//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateValues_FromSearchParameters(expression *string, populateFrom *string) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	if populateFrom == nil {
		return fmt.Errorf("parameter populateFrom is required, but nil was provided")
	}

	return nil
}

func validateValues_FromSearchComponentsParameters(components *SearchComponents) error {
	if components == nil {
		return fmt.Errorf("parameter components is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(components, func() string { return "parameter components" }); err != nil {
		return err
	}

	return nil
}

func validateValues_FromValuesParameters(values *[]*VariableValue) error {
	for idx_89445e, v := range *values {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter values[%#v]", idx_89445e) }); err != nil {
			return err
		}
	}

	return nil
}

