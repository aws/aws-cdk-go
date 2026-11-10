//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebFunction) validateAddEndpointParameters(id *string, options *WebFunctionEndpointOptions) error {
	return nil
}

func (w *jsiiProxy_WebFunction) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (w *jsiiProxy_WebFunction) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (w *jsiiProxy_WebFunction) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (w *jsiiProxy_WebFunction) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WebFunction) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WebFunction) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateWebFunction_FromWebFunctionArnParameters(scope constructs.Construct, id *string, webFunctionArn *string) error {
	return nil
}

func validateWebFunction_FromWebFunctionAttributesParameters(scope constructs.Construct, id *string, attrs *WebFunctionAttributes) error {
	return nil
}

func validateWebFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWebFunction_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWebFunction_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWebFunction_IsWebFunctionParameters(x interface{}) error {
	return nil
}

func validateNewWebFunctionParameters(scope constructs.Construct, id *string, props *WebFunctionProps) error {
	return nil
}

