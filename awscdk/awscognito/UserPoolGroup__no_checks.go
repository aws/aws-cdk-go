//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPoolGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPoolGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateUserPoolGroup_FromGroupNameParameters(scope constructs.Construct, id *string, groupName *string) error {
	return nil
}

func validateUserPoolGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPoolGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUserPoolGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUserPoolGroupParameters(scope constructs.Construct, id *string, props *UserPoolGroupProps) error {
	return nil
}

