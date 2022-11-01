//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClientVpnAuthorizationRule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ClientVpnAuthorizationRule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ClientVpnAuthorizationRule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_ClientVpnAuthorizationRule) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_ClientVpnAuthorizationRule) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateClientVpnAuthorizationRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClientVpnAuthorizationRule_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewClientVpnAuthorizationRuleParameters(scope constructs.Construct, id *string, props *ClientVpnAuthorizationRuleProps) error {
	return nil
}

