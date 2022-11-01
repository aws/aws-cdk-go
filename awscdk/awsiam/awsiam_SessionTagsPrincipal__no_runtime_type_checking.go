//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SessionTagsPrincipal) validateAddToAssumeRolePolicyParameters(doc PolicyDocument) error {
	return nil
}

func (s *jsiiProxy_SessionTagsPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SessionTagsPrincipal) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SessionTagsPrincipal) validateAppendDedupeParameters(append *string) error {
	return nil
}

func (s *jsiiProxy_SessionTagsPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewSessionTagsPrincipalParameters(principal IPrincipal) error {
	return nil
}

