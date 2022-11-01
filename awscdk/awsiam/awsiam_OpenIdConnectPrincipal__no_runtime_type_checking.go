//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpenIdConnectPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	return nil
}

func (o *jsiiProxy_OpenIdConnectPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (o *jsiiProxy_OpenIdConnectPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	return nil
}

func (o *jsiiProxy_OpenIdConnectPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewOpenIdConnectPrincipalParameters(openIdConnectProvider IOpenIdConnectProvider) error {
	return nil
}

