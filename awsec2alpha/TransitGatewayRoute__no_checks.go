//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayRoute) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRoute) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRoute) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGatewayRoute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayRoute_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayRoute_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayRouteParameters(scope constructs.Construct, id *string, props *TransitGatewayRouteProps) error {
	return nil
}

