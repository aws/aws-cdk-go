//go:build no_runtime_type_checking

package awscdkpipesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InputTransformation) validateBindParameters(pipe IPipe) error {
	return nil
}

func validateInputTransformation_FromEventPathParameters(jsonPathExpression *string) error {
	return nil
}

func validateInputTransformation_FromObjectParameters(inputTemplate *map[string]interface{}) error {
	return nil
}

func validateInputTransformation_FromTextParameters(inputTemplate *string) error {
	return nil
}

