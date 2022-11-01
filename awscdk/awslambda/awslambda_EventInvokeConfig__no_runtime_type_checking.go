//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventInvokeConfig) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EventInvokeConfig) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EventInvokeConfig) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_EventInvokeConfig) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (e *jsiiProxy_EventInvokeConfig) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateEventInvokeConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEventInvokeConfig_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewEventInvokeConfigParameters(scope constructs.Construct, id *string, props *EventInvokeConfigProps) error {
	return nil
}

