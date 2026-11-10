//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebFunctionBase) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (w *jsiiProxy_WebFunctionBase) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (w *jsiiProxy_WebFunctionBase) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (w *jsiiProxy_WebFunctionBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WebFunctionBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WebFunctionBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateWebFunctionBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWebFunctionBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWebFunctionBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewWebFunctionBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

