//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClientVpnRoute) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ClientVpnRoute) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ClientVpnRoute) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateClientVpnRoute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClientVpnRoute_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateClientVpnRoute_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewClientVpnRouteParameters(scope constructs.Construct, id *string, props *ClientVpnRouteProps) error {
	return nil
}

