//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrincipalWithConditions) validateAddConditionParameters(key *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateAddConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateAppendDedupeParameters(append *string) error {
	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewPrincipalWithConditionsParameters(principal IPrincipal, conditions *map[string]interface{}) error {
	return nil
}

