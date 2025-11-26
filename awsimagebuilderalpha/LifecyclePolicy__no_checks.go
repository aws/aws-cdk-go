//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LifecyclePolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LifecyclePolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LifecyclePolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (l *jsiiProxy_LifecyclePolicy) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (l *jsiiProxy_LifecyclePolicy) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateLifecyclePolicy_FromLifecyclePolicyArnParameters(scope constructs.Construct, id *string, lifecyclePolicyArn *string) error {
	return nil
}

func validateLifecyclePolicy_FromLifecyclePolicyNameParameters(scope constructs.Construct, id *string, lifecyclePolicyName *string) error {
	return nil
}

func validateLifecyclePolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLifecyclePolicy_IsLifecyclePolicyParameters(x interface{}) error {
	return nil
}

func validateLifecyclePolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLifecyclePolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLifecyclePolicyParameters(scope constructs.Construct, id *string, props *LifecyclePolicyProps) error {
	return nil
}

