//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"
)

func (i *jsiiProxy_InputTransformation) validateBindParameters(pipe IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func validateInputTransformation_FromEventPathParameters(jsonPathExpression *string) error {
	if jsonPathExpression == nil {
		return fmt.Errorf("parameter jsonPathExpression is required, but nil was provided")
	}

	return nil
}

func validateInputTransformation_FromObjectParameters(inputTemplate *map[string]interface{}) error {
	if inputTemplate == nil {
		return fmt.Errorf("parameter inputTemplate is required, but nil was provided")
	}

	return nil
}

func validateInputTransformation_FromTextParameters(inputTemplate *string) error {
	if inputTemplate == nil {
		return fmt.Errorf("parameter inputTemplate is required, but nil was provided")
	}

	return nil
}

