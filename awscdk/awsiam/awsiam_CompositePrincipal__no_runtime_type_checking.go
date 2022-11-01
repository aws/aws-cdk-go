//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CompositePrincipal) validateAddToAssumeRolePolicyParameters(doc PolicyDocument) error {
	return nil
}

func (c *jsiiProxy_CompositePrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CompositePrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CompositePrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

