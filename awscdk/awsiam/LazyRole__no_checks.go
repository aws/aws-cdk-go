//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LazyRole) validateAddManagedPolicyParameters(policy IManagedPolicy) error {
	return nil
}

func (l *jsiiProxy_LazyRole) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (l *jsiiProxy_LazyRole) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (l *jsiiProxy_LazyRole) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LazyRole) validateAttachInlinePolicyParameters(policy Policy) error {
	return nil
}

func (l *jsiiProxy_LazyRole) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LazyRole) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (l *jsiiProxy_LazyRole) validateGrantParameters(identity IPrincipal) error {
	return nil
}

func (l *jsiiProxy_LazyRole) validateGrantAssumeRoleParameters(identity IPrincipal) error {
	return nil
}

func (l *jsiiProxy_LazyRole) validateGrantPassRoleParameters(identity IPrincipal) error {
	return nil
}

func validateLazyRole_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLazyRole_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLazyRole_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLazyRoleParameters(scope constructs.Construct, id *string, props *LazyRoleProps) error {
	return nil
}

