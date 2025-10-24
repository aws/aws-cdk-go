//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RuntimeEndpoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RuntimeEndpoint) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RuntimeEndpoint) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRuntimeEndpoint_FromRuntimeEndpointAttributesParameters(scope constructs.Construct, id *string, attrs *RuntimeEndpointAttributes) error {
	return nil
}

func validateRuntimeEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRuntimeEndpoint_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRuntimeEndpoint_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRuntimeEndpointParameters(scope constructs.Construct, id *string, props *RuntimeEndpointProps) error {
	return nil
}

