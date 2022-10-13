//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiMapping) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_ApiMapping) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_ApiMapping) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_ApiMapping) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_ApiMapping) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateApiMapping_FromApiMappingAttributesParameters(scope constructs.Construct, id *string, attrs *ApiMappingAttributes) error {
	return nil
}

func validateApiMapping_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApiMapping_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewApiMappingParameters(scope constructs.Construct, id *string, props *ApiMappingProps) error {
	return nil
}

