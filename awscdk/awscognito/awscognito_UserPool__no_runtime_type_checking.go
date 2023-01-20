//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPool) validateAddClientParameters(id *string, options *UserPoolClientOptions) error {
	return nil
}

func (u *jsiiProxy_UserPool) validateAddDomainParameters(id *string, options *UserPoolDomainOptions) error {
	return nil
}

func (u *jsiiProxy_UserPool) validateAddResourceServerParameters(id *string, options *UserPoolResourceServerOptions) error {
	return nil
}

func (u *jsiiProxy_UserPool) validateAddTriggerParameters(operation UserPoolOperation, fn awslambda.IFunction) error {
	return nil
}

func (u *jsiiProxy_UserPool) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPool) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPool) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (u *jsiiProxy_UserPool) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (u *jsiiProxy_UserPool) validateRegisterIdentityProviderParameters(provider IUserPoolIdentityProvider) error {
	return nil
}

func validateUserPool_FromUserPoolArnParameters(scope constructs.Construct, id *string, userPoolArn *string) error {
	return nil
}

func validateUserPool_FromUserPoolIdParameters(scope constructs.Construct, id *string, userPoolId *string) error {
	return nil
}

func validateUserPool_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPool_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUserPool_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUserPoolParameters(scope constructs.Construct, id *string, props *UserPoolProps) error {
	return nil
}

