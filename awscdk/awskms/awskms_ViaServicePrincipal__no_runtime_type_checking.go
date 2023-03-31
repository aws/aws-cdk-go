//go:build no_runtime_type_checking

package awskms

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_ViaServicePrincipal) validateAddToAssumeRolePolicyParameters(document awsiam.PolicyDocument) error {
	return nil
}

func (v *jsiiProxy_ViaServicePrincipal) validateAddToPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (v *jsiiProxy_ViaServicePrincipal) validateAddToPrincipalPolicyParameters(_statement awsiam.PolicyStatement) error {
	return nil
}

func (v *jsiiProxy_ViaServicePrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewViaServicePrincipalParameters(serviceName *string) error {
	return nil
}

