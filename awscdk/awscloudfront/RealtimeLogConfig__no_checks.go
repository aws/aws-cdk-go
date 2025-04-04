//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RealtimeLogConfig) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RealtimeLogConfig) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RealtimeLogConfig) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRealtimeLogConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRealtimeLogConfig_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRealtimeLogConfig_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRealtimeLogConfigParameters(scope constructs.Construct, id *string, props *RealtimeLogConfigProps) error {
	return nil
}

