//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GatewayRoute) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GatewayRoute) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GatewayRoute) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateGatewayRoute_FromGatewayRouteArnParameters(scope constructs.Construct, id *string, gatewayRouteArn *string) error {
	return nil
}

func validateGatewayRoute_FromGatewayRouteAttributesParameters(scope constructs.Construct, id *string, attrs *GatewayRouteAttributes) error {
	return nil
}

func validateGatewayRoute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGatewayRoute_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGatewayRoute_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGatewayRouteParameters(scope constructs.Construct, id *string, props *GatewayRouteProps) error {
	return nil
}

