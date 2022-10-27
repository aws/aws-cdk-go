//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebIdentityPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (w *jsiiProxy_WebIdentityPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (w *jsiiProxy_WebIdentityPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (w *jsiiProxy_WebIdentityPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewWebIdentityPrincipalParameters(identityProvider *string) error {
	return nil
}

