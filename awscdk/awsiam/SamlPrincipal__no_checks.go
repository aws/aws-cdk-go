//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SamlPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (s *jsiiProxy_SamlPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SamlPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SamlPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewSamlPrincipalParameters(samlProvider ISamlProvider, conditions *map[string]interface{}) error {
	return nil
}

