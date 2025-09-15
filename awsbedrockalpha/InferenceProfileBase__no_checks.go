//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InferenceProfileBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_InferenceProfileBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_InferenceProfileBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_InferenceProfileBase) validateGrantProfileUsageParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateInferenceProfileBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInferenceProfileBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateInferenceProfileBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewInferenceProfileBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

