//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCognitoUserPoolsAuthorizer_IsAuthorizerParameters(x interface{}) error {
	return nil
}

func validateCognitoUserPoolsAuthorizer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCognitoUserPoolsAuthorizer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCognitoUserPoolsAuthorizer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCognitoUserPoolsAuthorizerParameters(scope constructs.Construct, id *string, props *CognitoUserPoolsAuthorizerProps) error {
	return nil
}

