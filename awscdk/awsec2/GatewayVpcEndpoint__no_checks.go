//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GatewayVpcEndpoint) validateAddToPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (g *jsiiProxy_GatewayVpcEndpoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GatewayVpcEndpoint) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GatewayVpcEndpoint) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateGatewayVpcEndpoint_FromGatewayVpcEndpointIdParameters(scope constructs.Construct, id *string, gatewayVpcEndpointId *string) error {
	return nil
}

func validateGatewayVpcEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGatewayVpcEndpoint_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGatewayVpcEndpoint_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGatewayVpcEndpointParameters(scope constructs.Construct, id *string, props *GatewayVpcEndpointProps) error {
	return nil
}

