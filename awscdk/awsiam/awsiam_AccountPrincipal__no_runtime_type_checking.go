//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccountPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (a *jsiiProxy_AccountPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_AccountPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_AccountPrincipal) validateInOrganizationParameters(organizationId *string) error {
	return nil
}

func (a *jsiiProxy_AccountPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewAccountPrincipalParameters(accountId interface{}) error {
	return nil
}

