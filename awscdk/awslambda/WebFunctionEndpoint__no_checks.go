//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebFunctionEndpoint) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (w *jsiiProxy_WebFunctionEndpoint) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (w *jsiiProxy_WebFunctionEndpoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WebFunctionEndpoint) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WebFunctionEndpoint) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateWebFunctionEndpoint_FromWebFunctionEndpointAttributesParameters(scope constructs.Construct, id *string, attrs *WebFunctionEndpointAttributes) error {
	return nil
}

func validateWebFunctionEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWebFunctionEndpoint_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWebFunctionEndpoint_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWebFunctionEndpoint_IsWebFunctionEndpointParameters(x interface{}) error {
	return nil
}

func validateNewWebFunctionEndpointParameters(scope constructs.Construct, id *string, props *WebFunctionEndpointProps) error {
	return nil
}

