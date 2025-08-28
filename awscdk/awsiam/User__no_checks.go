//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_User) validateAddManagedPolicyParameters(policy IManagedPolicy) error {
	return nil
}

func (u *jsiiProxy_User) validateAddToGroupParameters(group IGroup) error {
	return nil
}

func (u *jsiiProxy_User) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (u *jsiiProxy_User) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (u *jsiiProxy_User) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_User) validateAttachInlinePolicyParameters(policy Policy) error {
	return nil
}

func (u *jsiiProxy_User) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_User) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateUser_FromUserArnParameters(scope constructs.Construct, id *string, userArn *string) error {
	return nil
}

func validateUser_FromUserAttributesParameters(scope constructs.Construct, id *string, attrs *UserAttributes) error {
	return nil
}

func validateUser_FromUserNameParameters(scope constructs.Construct, id *string, userName *string) error {
	return nil
}

func validateUser_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUser_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUser_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUserParameters(scope constructs.Construct, id *string, props *UserProps) error {
	return nil
}

