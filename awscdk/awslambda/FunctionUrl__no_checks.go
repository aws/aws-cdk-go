//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionUrl) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FunctionUrl) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FunctionUrl) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_FunctionUrl) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateFunctionUrl_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFunctionUrl_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFunctionUrl_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFunctionUrlParameters(scope constructs.Construct, id *string, props *FunctionUrlProps) error {
	return nil
}

