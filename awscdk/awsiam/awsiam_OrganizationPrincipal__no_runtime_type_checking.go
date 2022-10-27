//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OrganizationPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (o *jsiiProxy_OrganizationPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (o *jsiiProxy_OrganizationPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (o *jsiiProxy_OrganizationPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewOrganizationPrincipalParameters(organizationId *string) error {
	return nil
}

