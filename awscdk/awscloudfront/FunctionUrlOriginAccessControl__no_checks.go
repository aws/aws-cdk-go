//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionUrlOriginAccessControl) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FunctionUrlOriginAccessControl) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FunctionUrlOriginAccessControl) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateFunctionUrlOriginAccessControl_FromOriginAccessControlIdParameters(scope constructs.Construct, id *string, originAccessControlId *string) error {
	return nil
}

func validateFunctionUrlOriginAccessControl_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFunctionUrlOriginAccessControl_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFunctionUrlOriginAccessControl_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFunctionUrlOriginAccessControlParameters(scope constructs.Construct, id *string, props *FunctionUrlOriginAccessControlProps) error {
	return nil
}

