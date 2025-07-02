//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UnmanagedComputeEnvironment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UnmanagedComputeEnvironment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UnmanagedComputeEnvironment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateUnmanagedComputeEnvironment_FromUnmanagedComputeEnvironmentArnParameters(scope constructs.Construct, id *string, unmanagedComputeEnvironmentArn *string) error {
	return nil
}

func validateUnmanagedComputeEnvironment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUnmanagedComputeEnvironment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUnmanagedComputeEnvironment_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUnmanagedComputeEnvironmentParameters(scope constructs.Construct, id *string, props *UnmanagedComputeEnvironmentProps) error {
	return nil
}

