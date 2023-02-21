//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IIdentity) validateAddManagedPolicyParameters(policy IManagedPolicy) error {
	return nil
}

func (i *jsiiProxy_IIdentity) validateAttachInlinePolicyParameters(policy Policy) error {
	return nil
}

func (i *jsiiProxy_IIdentity) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IIdentity) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

