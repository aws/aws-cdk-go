//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (m *jsiiProxy_ManagedPolicy) validateAttachToGroupParameters(group IGroup) error {
	return nil
}

func (m *jsiiProxy_ManagedPolicy) validateAttachToRoleParameters(role IRole) error {
	return nil
}

func (m *jsiiProxy_ManagedPolicy) validateAttachToUserParameters(user IUser) error {
	return nil
}

func (m *jsiiProxy_ManagedPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (m *jsiiProxy_ManagedPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateManagedPolicy_FromAwsManagedPolicyNameParameters(managedPolicyName *string) error {
	return nil
}

func validateManagedPolicy_FromManagedPolicyArnParameters(scope constructs.Construct, id *string, managedPolicyArn *string) error {
	return nil
}

func validateManagedPolicy_FromManagedPolicyNameParameters(scope constructs.Construct, id *string, managedPolicyName *string) error {
	return nil
}

func validateManagedPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateManagedPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateManagedPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewManagedPolicyParameters(scope constructs.Construct, id *string, props *ManagedPolicyProps) error {
	return nil
}

