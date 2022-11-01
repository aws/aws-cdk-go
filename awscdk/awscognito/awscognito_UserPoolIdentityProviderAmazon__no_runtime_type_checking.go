//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderAmazon) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateUserPoolIdentityProviderAmazon_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPoolIdentityProviderAmazon_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewUserPoolIdentityProviderAmazonParameters(scope constructs.Construct, id *string, props *UserPoolIdentityProviderAmazonProps) error {
	return nil
}

