//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Policy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_Policy) validateAttachToGroupParameters(group IGroup) error {
	return nil
}

func (p *jsiiProxy_Policy) validateAttachToRoleParameters(role IRole) error {
	return nil
}

func (p *jsiiProxy_Policy) validateAttachToUserParameters(user IUser) error {
	return nil
}

func (p *jsiiProxy_Policy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_Policy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePolicy_FromPolicyNameParameters(scope constructs.Construct, id *string, policyName *string) error {
	return nil
}

func validatePolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPolicyParameters(scope constructs.Construct, id *string, props *PolicyProps) error {
	return nil
}

