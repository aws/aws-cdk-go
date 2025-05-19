//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CanonicalUserPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (c *jsiiProxy_CanonicalUserPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CanonicalUserPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CanonicalUserPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewCanonicalUserPrincipalParameters(canonicalUserId *string) error {
	return nil
}

