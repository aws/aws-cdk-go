//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IUser) validateAddToGroupParameters(group IGroup) error {
	return nil
}

func (i *jsiiProxy_IUser) validateAddManagedPolicyParameters(policy IManagedPolicy) error {
	return nil
}

func (i *jsiiProxy_IUser) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IUser) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IUser) validateAttachInlinePolicyParameters(policy Policy) error {
	return nil
}

