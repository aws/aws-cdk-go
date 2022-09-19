//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IdentityPool) validateAddRoleMappingsParameters(roleMappings *[]*IdentityPoolRoleMapping) error {
	return nil
}

func (i *jsiiProxy_IdentityPool) validateAddUserPoolAuthenticationParameters(userPool IUserPoolAuthenticationProvider) error {
	return nil
}

func (i *jsiiProxy_IdentityPool) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IdentityPool) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IdentityPool) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIdentityPool_FromIdentityPoolArnParameters(scope constructs.Construct, id *string, identityPoolArn *string) error {
	return nil
}

func validateIdentityPool_FromIdentityPoolIdParameters(scope constructs.Construct, id *string, identityPoolId *string) error {
	return nil
}

func validateIdentityPool_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIdentityPool_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIdentityPool_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIdentityPoolParameters(scope constructs.Construct, id *string, props *IdentityPoolProps) error {
	return nil
}

