//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SamlConsolePrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (s *jsiiProxy_SamlConsolePrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SamlConsolePrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SamlConsolePrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewSamlConsolePrincipalParameters(samlProvider ISamlProvider) error {
	return nil
}

