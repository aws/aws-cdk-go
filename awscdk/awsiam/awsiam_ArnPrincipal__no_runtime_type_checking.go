//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ArnPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (a *jsiiProxy_ArnPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_ArnPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_ArnPrincipal) validateInOrganizationParameters(organizationId *string) error {
	return nil
}

func (a *jsiiProxy_ArnPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewArnPrincipalParameters(arn *string) error {
	return nil
}

