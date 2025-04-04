//go:build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpIntegration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (h *jsiiProxy_HttpIntegration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (h *jsiiProxy_HttpIntegration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateHttpIntegration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHttpIntegration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateHttpIntegration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewHttpIntegrationParameters(scope constructs.Construct, id *string, props *HttpIntegrationProps) error {
	return nil
}

