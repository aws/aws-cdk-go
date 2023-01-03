//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolIdentityProviderOidc) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderOidc) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderOidc) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderOidc) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderOidc) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateUserPoolIdentityProviderOidc_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPoolIdentityProviderOidc_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewUserPoolIdentityProviderOidcParameters(scope constructs.Construct, id *string, props *UserPoolIdentityProviderOidcProps) error {
	return nil
}

