//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LifecycleHook) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LifecycleHook) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LifecycleHook) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateLifecycleHook_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLifecycleHook_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLifecycleHook_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLifecycleHookParameters(scope constructs.Construct, id *string, props *LifecycleHookProps) error {
	return nil
}

