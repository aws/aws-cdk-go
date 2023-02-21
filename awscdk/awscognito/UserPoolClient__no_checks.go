//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolClient) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPoolClient) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPoolClient) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateUserPoolClient_FromUserPoolClientIdParameters(scope constructs.Construct, id *string, userPoolClientId *string) error {
	return nil
}

func validateUserPoolClient_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPoolClient_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUserPoolClient_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUserPoolClientParameters(scope constructs.Construct, id *string, props *UserPoolClientProps) error {
	return nil
}

