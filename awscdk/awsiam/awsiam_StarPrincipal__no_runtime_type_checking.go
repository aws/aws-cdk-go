//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StarPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (s *jsiiProxy_StarPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_StarPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_StarPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

