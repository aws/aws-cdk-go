//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Anyone) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (a *jsiiProxy_Anyone) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_Anyone) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_Anyone) validateInOrganizationParameters(organizationId *string) error {
	return nil
}

func (a *jsiiProxy_Anyone) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

