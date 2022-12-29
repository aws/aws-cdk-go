//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClientVpnEndpoint) validateAddAuthorizationRuleParameters(id *string, props *ClientVpnAuthorizationRuleOptions) error {
	return nil
}

func (c *jsiiProxy_ClientVpnEndpoint) validateAddRouteParameters(id *string, props *ClientVpnRouteOptions) error {
	return nil
}

func (c *jsiiProxy_ClientVpnEndpoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ClientVpnEndpoint) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ClientVpnEndpoint) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateClientVpnEndpoint_FromEndpointAttributesParameters(scope constructs.Construct, id *string, attrs *ClientVpnEndpointAttributes) error {
	return nil
}

func validateClientVpnEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClientVpnEndpoint_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateClientVpnEndpoint_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewClientVpnEndpointParameters(scope constructs.Construct, id *string, props *ClientVpnEndpointProps) error {
	return nil
}

