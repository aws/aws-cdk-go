//go:build no_runtime_type_checking

package awscdkelasticachealpha

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NoPasswordUser) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NoPasswordUser) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NoPasswordUser) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNoPasswordUser_FromUserArnParameters(scope constructs.Construct, id *string, userArn *string) error {
	return nil
}

func validateNoPasswordUser_FromUserAttributesParameters(scope constructs.Construct, id *string, attrs *UserBaseAttributes) error {
	return nil
}

func validateNoPasswordUser_FromUserIdParameters(scope constructs.Construct, id *string, userId *string) error {
	return nil
}

func validateNoPasswordUser_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNoPasswordUser_IsNoPasswordUserParameters(x interface{}) error {
	return nil
}

func validateNoPasswordUser_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNoPasswordUser_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNoPasswordUserParameters(scope constructs.Construct, id *string, props *NoPasswordUserProps) error {
	return nil
}

