//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IGroup) validateAddManagedPolicyParameters(policy IManagedPolicy) error {
	return nil
}

func (i *jsiiProxy_IGroup) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IGroup) validateAttachInlinePolicyParameters(policy Policy) error {
	return nil
}

