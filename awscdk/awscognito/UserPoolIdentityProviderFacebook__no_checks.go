//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderFacebook) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateUserPoolIdentityProviderFacebook_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPoolIdentityProviderFacebook_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUserPoolIdentityProviderFacebook_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUserPoolIdentityProviderFacebookParameters(scope constructs.Construct, id *string, props *UserPoolIdentityProviderFacebookProps) error {
	return nil
}

