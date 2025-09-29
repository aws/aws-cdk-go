//go:build no_runtime_type_checking

package awscdkelasticachealpha

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserGroup) validateAddUserParameters(user IUser) error {
	return nil
}

func (u *jsiiProxy_UserGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateUserGroup_FromUserGroupArnParameters(scope constructs.Construct, id *string, userGroupArn *string) error {
	return nil
}

func validateUserGroup_FromUserGroupAttributesParameters(scope constructs.Construct, id *string, attrs *UserGroupAttributes) error {
	return nil
}

func validateUserGroup_FromUserGroupNameParameters(scope constructs.Construct, id *string, userGroupName *string) error {
	return nil
}

func validateUserGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUserGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUserGroup_IsUserGroupParameters(x interface{}) error {
	return nil
}

func validateNewUserGroupParameters(scope constructs.Construct, id *string, props *UserGroupProps) error {
	return nil
}

