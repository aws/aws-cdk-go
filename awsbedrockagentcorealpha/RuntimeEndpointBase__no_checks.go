//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RuntimeEndpointBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RuntimeEndpointBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RuntimeEndpointBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRuntimeEndpointBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRuntimeEndpointBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRuntimeEndpointBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRuntimeEndpointBaseParameters(scope constructs.Construct, id *string) error {
	return nil
}

