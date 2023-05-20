//go:build no_runtime_type_checking

package awscdkapigatewayv2alpha

// Building without runtime type checking enabled, so all the below just return nil

func validateMappingValue_ContextVariableParameters(variableName *string) error {
	return nil
}

func validateMappingValue_CustomParameters(value *string) error {
	return nil
}

func validateMappingValue_RequestBodyParameters(name *string) error {
	return nil
}

func validateMappingValue_RequestHeaderParameters(name *string) error {
	return nil
}

func validateMappingValue_RequestPathParamParameters(name *string) error {
	return nil
}

func validateMappingValue_RequestQueryStringParameters(name *string) error {
	return nil
}

func validateMappingValue_StageVariableParameters(variableName *string) error {
	return nil
}

func validateNewMappingValueParameters(value *string) error {
	return nil
}

