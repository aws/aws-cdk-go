//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccountRootPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (a *jsiiProxy_AccountRootPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_AccountRootPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_AccountRootPrincipal) validateInOrganizationParameters(organizationId *string) error {
	return nil
}

func (a *jsiiProxy_AccountRootPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

