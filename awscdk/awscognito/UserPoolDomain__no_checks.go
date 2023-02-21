//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolDomain) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPoolDomain) validateBaseUrlParameters(options *BaseUrlOptions) error {
	return nil
}

func (u *jsiiProxy_UserPoolDomain) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPoolDomain) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (u *jsiiProxy_UserPoolDomain) validateSignInUrlParameters(client UserPoolClient, options *SignInUrlOptions) error {
	return nil
}

func validateUserPoolDomain_FromDomainNameParameters(scope constructs.Construct, id *string, userPoolDomainName *string) error {
	return nil
}

func validateUserPoolDomain_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPoolDomain_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUserPoolDomain_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUserPoolDomainParameters(scope constructs.Construct, id *string, props *UserPoolDomainProps) error {
	return nil
}

