//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_Function) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_Function) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_Function) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateFunction_FromFunctionAttributesParameters(scope constructs.Construct, id *string, attrs *FunctionAttributes) error {
	return nil
}

func validateFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFunction_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFunction_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFunctionParameters(scope constructs.Construct, id *string, props *FunctionProps) error {
	return nil
}

