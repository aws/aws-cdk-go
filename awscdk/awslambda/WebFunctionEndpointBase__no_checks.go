//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebFunctionEndpointBase) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (w *jsiiProxy_WebFunctionEndpointBase) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (w *jsiiProxy_WebFunctionEndpointBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WebFunctionEndpointBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WebFunctionEndpointBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateWebFunctionEndpointBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWebFunctionEndpointBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWebFunctionEndpointBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewWebFunctionEndpointBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

