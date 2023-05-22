//go:build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ParameterMapping) validateAppendHeaderParameters(name *string, value MappingValue) error {
	return nil
}

func (p *jsiiProxy_ParameterMapping) validateAppendQueryStringParameters(name *string, value MappingValue) error {
	return nil
}

func (p *jsiiProxy_ParameterMapping) validateCustomParameters(key *string, value *string) error {
	return nil
}

func (p *jsiiProxy_ParameterMapping) validateOverwriteHeaderParameters(name *string, value MappingValue) error {
	return nil
}

func (p *jsiiProxy_ParameterMapping) validateOverwritePathParameters(value MappingValue) error {
	return nil
}

func (p *jsiiProxy_ParameterMapping) validateOverwriteQueryStringParameters(name *string, value MappingValue) error {
	return nil
}

func (p *jsiiProxy_ParameterMapping) validateRemoveHeaderParameters(name *string) error {
	return nil
}

func (p *jsiiProxy_ParameterMapping) validateRemoveQueryStringParameters(name *string) error {
	return nil
}

func validateParameterMapping_FromObjectParameters(obj *map[string]MappingValue) error {
	return nil
}

