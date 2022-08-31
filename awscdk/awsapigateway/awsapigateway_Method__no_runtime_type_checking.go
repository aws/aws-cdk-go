//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Method) validateAddMethodResponseParameters(methodResponse *MethodResponse) error {
	return nil
}

func (m *jsiiProxy_Method) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (m *jsiiProxy_Method) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (m *jsiiProxy_Method) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (m *jsiiProxy_Method) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (m *jsiiProxy_Method) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateMethod_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMethod_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewMethodParameters(scope constructs.Construct, id *string, props *MethodProps) error {
	return nil
}

