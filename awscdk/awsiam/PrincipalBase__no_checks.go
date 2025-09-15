//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrincipalBase) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (p *jsiiProxy_PrincipalBase) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (p *jsiiProxy_PrincipalBase) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (p *jsiiProxy_PrincipalBase) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

