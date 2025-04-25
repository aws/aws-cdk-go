//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FederatedPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (f *jsiiProxy_FederatedPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (f *jsiiProxy_FederatedPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (f *jsiiProxy_FederatedPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewFederatedPrincipalParameters(federated *string) error {
	return nil
}

