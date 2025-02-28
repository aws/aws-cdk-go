//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AnyPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (a *jsiiProxy_AnyPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_AnyPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_AnyPrincipal) validateInOrganizationParameters(organizationId *string) error {
	return nil
}

func (a *jsiiProxy_AnyPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

