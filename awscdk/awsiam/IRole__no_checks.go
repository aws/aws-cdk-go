//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IRole) validateGrantParameters(grantee IPrincipal) error {
	return nil
}

func (i *jsiiProxy_IRole) validateGrantAssumeRoleParameters(grantee IPrincipal) error {
	return nil
}

func (i *jsiiProxy_IRole) validateGrantPassRoleParameters(grantee IPrincipal) error {
	return nil
}

func (i *jsiiProxy_IRole) validateAddManagedPolicyParameters(policy IManagedPolicy) error {
	return nil
}

func (i *jsiiProxy_IRole) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IRole) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IRole) validateAttachInlinePolicyParameters(policy Policy) error {
	return nil
}

