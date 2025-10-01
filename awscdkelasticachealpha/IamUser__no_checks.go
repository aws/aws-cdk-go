//go:build no_runtime_type_checking

package awscdkelasticachealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IamUser) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IamUser) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IamUser) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_IamUser) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IamUser) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateIamUser_FromUserArnParameters(scope constructs.Construct, id *string, userArn *string) error {
	return nil
}

func validateIamUser_FromUserAttributesParameters(scope constructs.Construct, id *string, attrs *UserBaseAttributes) error {
	return nil
}

func validateIamUser_FromUserIdParameters(scope constructs.Construct, id *string, userId *string) error {
	return nil
}

func validateIamUser_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIamUser_IsIamUserParameters(x interface{}) error {
	return nil
}

func validateIamUser_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIamUser_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIamUserParameters(scope constructs.Construct, id *string, props *IamUserProps) error {
	return nil
}

