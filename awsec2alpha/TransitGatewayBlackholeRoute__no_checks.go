//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayBlackholeRoute) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayBlackholeRoute) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayBlackholeRoute) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGatewayBlackholeRoute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayBlackholeRoute_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayBlackholeRoute_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayBlackholeRouteParameters(scope constructs.Construct, id *string, props *TransitGatewayBlackholeRouteProps) error {
	return nil
}

