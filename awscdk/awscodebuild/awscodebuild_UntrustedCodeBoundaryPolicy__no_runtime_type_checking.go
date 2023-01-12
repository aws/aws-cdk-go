//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateAttachToGroupParameters(group awsiam.IGroup) error {
	return nil
}

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateAttachToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateAttachToUserParameters(user awsiam.IUser) error {
	return nil
}

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateUntrustedCodeBoundaryPolicy_FromAwsManagedPolicyNameParameters(managedPolicyName *string) error {
	return nil
}

func validateUntrustedCodeBoundaryPolicy_FromManagedPolicyArnParameters(scope constructs.Construct, id *string, managedPolicyArn *string) error {
	return nil
}

func validateUntrustedCodeBoundaryPolicy_FromManagedPolicyNameParameters(scope constructs.Construct, id *string, managedPolicyName *string) error {
	return nil
}

func validateUntrustedCodeBoundaryPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUntrustedCodeBoundaryPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUntrustedCodeBoundaryPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUntrustedCodeBoundaryPolicyParameters(scope constructs.Construct, id *string, props *UntrustedCodeBoundaryPolicyProps) error {
	return nil
}

