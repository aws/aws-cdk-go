//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolIdentityProviderSaml) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderSaml) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderSaml) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateUserPoolIdentityProviderSaml_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPoolIdentityProviderSaml_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUserPoolIdentityProviderSaml_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUserPoolIdentityProviderSamlParameters(scope constructs.Construct, id *string, props *UserPoolIdentityProviderSamlProps) error {
	return nil
}

