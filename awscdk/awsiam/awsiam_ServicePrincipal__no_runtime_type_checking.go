//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServicePrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (s *jsiiProxy_ServicePrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_ServicePrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_ServicePrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateServicePrincipal_ServicePrincipalNameParameters(service *string) error {
	return nil
}

func validateNewServicePrincipalParameters(service *string, opts *ServicePrincipalOpts) error {
	return nil
}

