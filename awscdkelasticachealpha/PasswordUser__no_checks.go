//go:build no_runtime_type_checking

package awscdkelasticachealpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PasswordUser) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PasswordUser) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PasswordUser) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePasswordUser_FromUserArnParameters(scope constructs.Construct, id *string, userArn *string) error {
	return nil
}

func validatePasswordUser_FromUserAttributesParameters(scope constructs.Construct, id *string, attrs *UserBaseAttributes) error {
	return nil
}

func validatePasswordUser_FromUserIdParameters(scope constructs.Construct, id *string, userId *string) error {
	return nil
}

func validatePasswordUser_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePasswordUser_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePasswordUser_IsPasswordUserParameters(x interface{}) error {
	return nil
}

func validatePasswordUser_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPasswordUserParameters(scope constructs.Construct, id *string, props *PasswordUserProps) error {
	return nil
}

